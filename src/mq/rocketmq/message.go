package rocketmq

import "time"

type CacheMessage struct {
	Key       string `json:"key"`
	Value     any    `json:"value"`
	Operation string `json:"operation"`
	Timestamp int    `json:"timestamp"`
}

func NewCacheMessage(key string, value any, operation string) *CacheMessage {
	return &CacheMessage{
		Key:       key,
		Value:     value,
		Operation: operation,
		Timestamp: int(time.Now().Unix()),
	}
}
