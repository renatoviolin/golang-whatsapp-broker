package kafka

import (
	"time"
	"whatsapp-client/infra/logger"

	"github.com/Shopify/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer(servers []string) *Producer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Net.DialTimeout = 5 * time.Second

	producer, err := sarama.NewSyncProducer(servers, config)
	if err != nil {
		logger.Fatal("new-producer", err.Error())
	}

	return &Producer{producer: producer}
}

func (h *Producer) Produce(value []byte, topic string) (partition int32, offset int64, err error) {
	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}
	return h.producer.SendMessage(&msg)
}
