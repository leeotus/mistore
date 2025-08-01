package models

import (
	"context"
	"encoding/json"
	"mistore/src/db"
	"time"
)

type Session struct {
	ID   string
	Data map[string]any
}

const COOKIE_SESSNAME = "session_id"
const SESSION_PREFIX = "session:"
const SESSION_EXPIRE = 3600 // 1小时，单位秒

// @brief 保存session数据到Redis
func SetSession(ctx context.Context, sessionID string, data map[string]any) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return db.RedisDB.Set(ctx, sessionID, b, SESSION_EXPIRE*time.Second).Err()
}

// @brief 从Redis获取session数据
func GetSession(ctx context.Context, sessionID string) (map[string]any, error) {
	val, err := db.RedisDB.Get(ctx, sessionID).Result()
	if err != nil {
		return nil, err
	}
	var data map[string]any
	err = json.Unmarshal([]byte(val), &data)
	return data, err
}

// @brief 刷新session过期时间
func RefreshSession(ctx context.Context, sessionID string) error {
	return db.RedisDB.Expire(ctx, sessionID, SESSION_EXPIRE*time.Second).Err()
}

// @brief 删除session
func DeleteSession(ctx context.Context, sessionID string) error {
	return db.RedisDB.Del(ctx, sessionID).Err()
}

// @brief 刷新session过期时间（带最小刷新间隔，支持传入已获取的session数据）
// 推荐只保留本方法，避免重复定义
func RefreshSessionWithInterval(ctx context.Context, sessionID string, data map[string]any, minInterval time.Duration) error {
	if data == nil {
		val, err := db.RedisDB.Get(ctx, sessionID).Result()
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(val), &data)
		if err != nil {
			return err
		}
	}
	loginTimeVal, ok := data["login_time"]
	if !ok {
		return nil // 没有login_time字段，不刷新
	}
	var loginTime int64
	switch v := loginTimeVal.(type) {
	case float64:
		loginTime = int64(v)
	case int64:
		loginTime = v
	case int:
		loginTime = int64(v)
	}
	if TimeStamp()-loginTime >= int64(minInterval.Seconds()) {
		err := db.RedisDB.Expire(ctx, sessionID, SESSION_EXPIRE*time.Second).Err()
		if err == nil {
			data["login_time"] = TimeStamp()
			b, _ := json.Marshal(data)
			db.RedisDB.Set(ctx, sessionID, b, SESSION_EXPIRE*time.Second)
		}
	}
	return nil
}
