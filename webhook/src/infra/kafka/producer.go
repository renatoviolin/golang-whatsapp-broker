package kafka

import (
	"broker/infra/logger"
	"time"

	"github.com/Shopify/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewProducer(servers []string, topic string) *Producer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Net.DialTimeout = 5 * time.Second

	producer, err := sarama.NewSyncProducer(servers, config)
	if err != nil {
		logger.Fatal("kafka", "new-producer", err.Error())
	}

	return &Producer{producer: producer, topic: topic}
}

func (h *Producer) Produce(value []byte) (partition int32, offset int64, err error) {
	msg := sarama.ProducerMessage{
		Topic: h.topic,
		Value: sarama.ByteEncoder(value),
	}
	return h.producer.SendMessage(&msg)
}
