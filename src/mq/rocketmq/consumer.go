package rocketmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type Consumer struct {
	c rocketmq.PushConsumer
}

func NewConsumer(nameServer, topic, group string) (*Consumer, error) {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{nameServer}),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(group),
	)
	if err != nil {
		return nil, fmt.Errorf("create consumer error: %v", err)
	}

	return &Consumer{c: c}, nil
}

func (c *Consumer) Start(topic string, handler func(*CacheMessage) error) error {
	err := c.c.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, me ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range me {
			var cacheMsg CacheMessage
			if err1 := json.Unmarshal(msg.Body, &cacheMsg); err1 != nil {
				log.Printf("unmarshal message failed: %v", err1)
				return consumer.ConsumeRetryLater, nil
			}

			if err2 := handler(&cacheMsg); err2 != nil {
				log.Printf("handle message error: %v", err2)
				return consumer.ConsumeRetryLater, nil
			}
		}

		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		return fmt.Errorf("subscribe error: %v", err)
	}

	return c.c.Start()
}

func (c *Consumer) Shutdown() error {
	return c.c.Shutdown()
}
