package utils

// @brief Pool简单的协程池（线程池）实现
// @note 使用 GetPool(workerCount) 获取单例
// @note 使用 Submit(task func()) 提交任务
// @todo 后面还需要改进

import "sync"

type Pool struct {
	taskQueue chan func()
}

var (
	instance *Pool
	once     sync.Once
)

var TaskPool *Pool

const MAX_THREADS_NUM = 12

func init() {
	TaskPool = GetPool(MAX_THREADS_NUM)
}

// GetPool 获取协程池单例，workerCount 只在首次调用时生效
func GetPool(workerCount int) *Pool {
	once.Do(func() {
		instance = &Pool{
			taskQueue: make(chan func(), 100),
		}
		for i := 0; i < workerCount; i++ {
			go instance.worker()
		}
	})
	return instance
}

func (p *Pool) worker() {
	for task := range p.taskQueue {
		task()
	}
}

// Submit 向线程池提交一个任务
func (p *Pool) Submit(task func()) {
	p.taskQueue <- task
}
