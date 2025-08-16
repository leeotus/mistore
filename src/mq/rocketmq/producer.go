package rocketmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type Producer struct {
	p rocketmq.Producer
}

func NewProducer(nameServer string, topic string) (*Producer, error) {
	p, err1 := rocketmq.NewProducer(
		producer.WithNameServer([]string{nameServer}),
		producer.WithRetry(2),
		producer.WithGroupName("GID_CACHE_CONSISTENCY"),
	)
	if err1 != nil {
		return nil, fmt.Errorf("create producer error: %v", err1)
	}

	if err2 := p.Start(); err2 != nil {
		return nil, fmt.Errorf("start producer error: %v", err2)
	}

	return &Producer{p: p}, nil
}

func (p *Producer) SendCacheMessage(ctx context.Context, topic string, msg *CacheMessage) error {
	body, err1 := json.Marshal(msg)
	if err1 != nil {
		return fmt.Errorf("marshal message error: %v", err1)
	}
	result, err2 := p.p.SendSync(ctx, &primitive.Message{
		Topic: topic,
		Body:  body,
	})
	if err2 != nil {
		return fmt.Errorf("send message error: %v", err2)
	}
	log.Printf("send message success, result:: %v", result)
	return nil
}

func (p *Producer) Shutdown() error {
	return p.p.Shutdown()
}
