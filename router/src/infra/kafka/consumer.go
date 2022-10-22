package kafka

import (
	"broker/infra/logger"
	"context"
	"time"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	handler  sarama.ConsumerGroupHandler
	consumer sarama.ConsumerGroup
	topic    []string
}

func NewConsumer(servers []string, topic string, consumerGroup string, handler sarama.ConsumerGroupHandler) *Consumer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Net.DialTimeout = 10 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Rebalance.Timeout = 1 * time.Second
	consumer, err := sarama.NewConsumerGroup(servers, consumerGroup, config)
	if err != nil {
		logger.Fatal("new-consumer", err.Error())
	}
	logger.Info("new-consumer", "ready to consume topic: "+topic)
	return &Consumer{
		consumer: consumer,
		handler:  handler,
		topic:    []string{topic},
	}
}

func (h *Consumer) StartConsumer() {
	go func() {
		for {
			err := h.consumer.Consume(context.Background(), h.topic, h.handler)
			if err != nil {
				logger.Error("consumer-goroutine", err.Error())
			}
		}
	}()
}
