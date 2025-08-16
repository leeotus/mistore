package db

/**
 * @brief Redis数据库以及缓存相关的函数
 * @author leeotus
 * @email leeotus@163.com
 * @note 实现多级缓存架构: 应用层->本地缓存(L1)->Redis缓存(L2)->MySQL数据库
 * @note 1.使用LRU算法来淘汰本地"老旧的"缓存; 2.使用定时器+TTL来定时清理过期缓存
 */

import (
	"context"
	"encoding/json"
	"fmt"
	"mistore/src/mq/rocketmq"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

/**
 * @brief Redis线程池
 */

var RedisDB *redis.Client
var ctx = context.Background()

const (
	DEFAULT_LOCAL_CACHE_EXPIRATION = time.Minute * 2
	DEFAULT_REDIS_CACHE_EXPIRATION = time.Hour * 2
	DEFAULT_LRU_CAPACITY           = 10000
	DEFAULT_CLEANUP_TIMERTICK      = 60 * time.Second
	CACHE_TOPIC                    = "cache_topic"
	CACHE_GROUP_NAME               = "GID_CACHE_CONSISTENCY"
)

type CacheItem struct {
	Data       any
	Expiration time.Time
}

// @brief LRU节点
type LRUNode struct {
	Key        string
	Value      *CacheItem
	Prev, Next *LRUNode
}

type RDBCache struct {
	// @note todo 这里直接采用了sync.Map, 实际应用中可能会遇到性能瓶颈问题,需考虑其他组件,这里为实现功能而直接使用
	localCache sync.Map // 本地缓存 @note: sync.Map是一个并发安全的键值对存储结构

	capacity int                // 最大缓存容量
	size     int                // 当前缓存大小
	head     *LRUNode           // LRU链表头结点
	tail     *LRUNode           // LRU链表尾节点
	mutex    sync.RWMutex       // 用于保护LRU链表操作的锁
	producer *rocketmq.Producer // rocketmq
}

func NewRDBCache(capacity int) *RDBCache {
	if capacity <= 0 {
		capacity = DEFAULT_LRU_CAPACITY
	}

	// 创建哨兵节点
	head := &LRUNode{}
	tail := &LRUNode{}
	head.Next = tail
	tail.Prev = head

	return &RDBCache{
		capacity: capacity,
		size:     0,
		head:     head,
		tail:     tail,
	}
}

var RedisCache *RDBCache

func RedisCoreInit(addr, port, pwd string, db, maxpool, minIdleConns int, maxConnAge, idleTimeout time.Duration) {
	host := fmt.Sprintf("%s:%s", addr, port)
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     host, // Redis地址
		Password: pwd,  // 密码
		DB:       db,   // 默认采用的数据库编号

		// 连接池配置
		PoolSize:     maxpool,      // 最大活跃连接数
		MinIdleConns: minIdleConns, // 最小空闲连接数
		MaxConnAge:   maxConnAge,   // 连接的最大存活时间
		IdleTimeout:  idleTimeout,  // 空闲连接的超时时间
	})

	// 测试连接
	// ctx := context.Background()
	_, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis连接失败: %v", err))
	}
	// 初始化LRU缓存
	RedisCache = NewRDBCache(DEFAULT_LRU_CAPACITY)

	// 初始化rocketmq生产者
	// @note 暂时把生成者作为成员放在RDBCache结构里,当然也可以提取出来
	producer, err2 := rocketmq.NewProducer("127.0.0.1:9876", CACHE_TOPIC)
	if err2 != nil {
		panic(fmt.Sprintf("MQ生产者初始化失败: %v", err2))
	}
	RedisCache.producer = producer
	// @usage: 如何使用生产者:
	// msg := rocketmq.NewCacheMessage(key, value, "UPDATE")
	// RedisCache.producer.SendCacheMessage(ctx, CACHE_TOPIC, msg)

	// 消费者: 负责将Redis保持与MySQL的一致性,同时由于使用了本地缓存,本地缓存还需要和redis保持一致性
	consumer, err3 := rocketmq.NewConsumer("127.0.0.1:9876", CACHE_TOPIC, CACHE_GROUP_NAME)
	if err3 != nil {
		panic(fmt.Sprintf("MQ消费者初始化失败: %v", err3))
	}

	// 启动消费者线程
	go func() {
		consumer.Start(CACHE_TOPIC, func(msg *rocketmq.CacheMessage) error {
			switch msg.Operation {
			case "UPDATE":
				var data any
				json.Unmarshal(msg.Value.([]byte), &data)
				// 更新redis
				RedisCache.SetMultiLevel(ctx, msg.Key, data, DEFAULT_REDIS_CACHE_EXPIRATION, DEFAULT_LOCAL_CACHE_EXPIRATION)
			case "DELETE":
				RedisDB.Del(ctx, msg.Key)
				if value, ok1 := RedisCache.localCache.Load(msg.Key); ok1 {
					if node, ok := value.(*LRUNode); ok {
						RedisCache.removeFromCache(msg.Key, node)
					}
				}
			}
			return nil
		})
	}()

	// 启动本地缓存清理任务
	go RedisCache.startCleanupTask()
}

// @brief 检查本地缓存是否过期
func (r *RDBCache) isExpired(item *CacheItem) bool {
	return time.Now().After(item.Expiration)
}

/**------------------------------------------------------------------------
 *!                           Redis数据库操作
 *------------------------------------------------------------------------**/
/**
 * @brief 向Redis里设置数据
 * @param key 键
 * @param value 键对应的值
 * @param expiration 该键值对的过期时间
 */
func (r *RDBCache) Set(key string, value any, expiration int) {
	bytes, _ := json.Marshal(value)
	RedisDB.Set(ctx, key, string(bytes), time.Second*time.Duration(expiration))
}

/**
 * @brief Set函数,带context版
 */
func (r *RDBCache) SetWithContext(ic context.Context, key string, value any, expiration int) {
	bytes, _ := json.Marshal(value)
	RedisDB.Set(ic, key, string(bytes), time.Second*time.Duration(expiration))
}

/**
 * @brief 企图获取保存在Redis里的数据,以输入"key"作为键值
 * @param obj 若获取到数据,则保存在obj中
 * @return bool 成功获取数据返回true,否则返回false
 */
func (r *RDBCache) Get(key string, obj any) bool {
	redisStr, err1 := RedisDB.Get(ctx, key).Result()
	if err1 == nil && redisStr != "" {
		err2 := json.Unmarshal([]byte(redisStr), obj)
		return err2 == nil
	}
	return false
}

/**
 * @brief Get函数,带context版
 */
func (r *RDBCache) GetWithCache(ic context.Context, key string, obj any) bool {
	redisStr, err1 := RedisDB.Get(ic, key).Result()
	if err1 == nil && redisStr != "" {
		err2 := json.Unmarshal([]byte(redisStr), obj)
		return err2 == nil
	}
	return false
}

/**------------------------------------------------------------------------
 *!                           多级缓存
 *------------------------------------------------------------------------**/
// @brief 获取多级缓存里的数据
func (r *RDBCache) GetMultiLevel(ctx context.Context, key string, result any) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// L1: 查询本地缓存
	if value, ok := r.localCache.Load(key); ok {
		switch v := value.(type) {
		case *LRUNode:
			if !r.isExpired(v.Value) {
				r.moveToHead(v)
				return json.Unmarshal(v.Value.Data.([]byte), result)
			}
			r.removeFromCache(key, v)
		}
	}

	// L2: 如果没有在本地缓存中找到数据,则需要去Redis中查询
	val, err := RedisDB.Get(ctx, key).Result()
	if err == nil && val != "" {
		// Redis命中,需要写入本地缓存
		// r.setLocalCache(key, []byte(val), DEFAULT_LOCAL_CACHE_EXPIRATION)
		r.setLocalCacheLRU(key, []byte(val), DEFAULT_LOCAL_CACHE_EXPIRATION)
		return json.Unmarshal([]byte(val), result)
	}
	return fmt.Errorf("cache miss")
}

func (r *RDBCache) SetMultiLevel(ctx context.Context, key string, value any, redisTTL, localTTL time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// 写入Redis
	err = RedisDB.Set(ctx, key, data, redisTTL).Err()
	if err != nil {
		return err
	}

	// 写入本地缓存
	r.setLocalCacheLRU(key, data, localTTL)
	return nil
}

/**------------------------------------------------------------------------
 *!                           定时器+TTL
 *------------------------------------------------------------------------**/

// @brief 定期清理过期的本地缓存项
func (r *RDBCache) cleanupExpired() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	// 从LRU链表尾部开始检查,限制扫描数量
	current := r.tail.Prev
	expiredCount := 0
	maxCleanup := 100 // 每次清理的最大个数
	for current != r.head && expiredCount < maxCleanup {
		if r.isExpired(current.Value) {
			prev := current.Prev
			r.removeFromCache(current.Key, current)
			expiredCount++
			current = prev
		} else {
			break // 由于是按照时间排序的,如果当前的没过期,前面的也不过期
		}
	}
}

func (r *RDBCache) startCleanupTask() {
	// 默认设置每30s清理一次过期缓存
	// 设置一个定时器
	ticker := time.NewTicker(DEFAULT_CLEANUP_TIMERTICK)
	defer ticker.Stop()

	for range ticker.C {
		r.cleanupExpired()
	}
}

/**------------------------------------------------------------------------
 *!                           LRU 算法
 *------------------------------------------------------------------------**/
// @brief 将新的LRU节点保存在LRU链表头部
// @todo 之后打算效仿MySQL或者Redis里的LRU算法来改进
func (r *RDBCache) addToHead(node *LRUNode) {
	node.Prev = r.head
	node.Next = r.head.Next
	r.head.Next.Prev = node
	r.head.Next = node
}

// @brief 删除一个LRU节点
func (r *RDBCache) removeNode(node *LRUNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// @brief 移动LRU节点
func (r *RDBCache) moveToHead(node *LRUNode) {
	r.removeNode(node)
	r.addToHead(node)
}

// @brief 移除尾部节点
func (r *RDBCache) removeTail() *LRUNode {
	lastNode := r.tail.Prev
	r.removeNode(lastNode)
	return lastNode
}

// @brief 设置LRU的方法
func (r *RDBCache) setLocalCacheLRU(key string, data []byte, ttl time.Duration) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	item := &CacheItem{
		Data:       data,
		Expiration: time.Now().Add(ttl),
	}
	if value, exist := r.localCache.Load(key); exist {
		// 更新节点
		node := value.(*LRUNode)
		node.Value = item
		r.moveToHead(node)
	} else {
		// 新增节点
		newNode := &LRUNode{
			Key:   key,
			Value: item,
		}

		if r.size >= r.capacity {
			// LRU淘汰
			tail := r.removeTail()
			r.localCache.Delete(tail.Key)
			r.size--
		}
		// 添加新节点
		r.addToHead(newNode)
		r.localCache.Store(key, newNode)
		r.size++
	}
}

// 从缓存中移除节点
func (r *RDBCache) removeFromCache(key string, node *LRUNode) {
	r.removeNode(node)
	r.localCache.Delete(key)
	r.size--
}
