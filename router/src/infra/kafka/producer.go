package kafka

import (
	"broker/dto"
	"broker/infra/logger"
	"encoding/json"
	"time"

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

func (h *Producer) Produce(value dto.KafkaOutputMessage, topic string) (partition int32, offset int64, err error) {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return 0, 0, err
	}
	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(valueBytes),
	}
	return h.producer.SendMessage(&msg)
}

func (h *Producer) ProduceError(value dto.KafkaErrorMessage, errorTopic string) (partition int32, offset int64, err error) {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return 0, 0, err
	}
	msg := sarama.ProducerMessage{
		Topic: errorTopic,
		Value: sarama.ByteEncoder(valueBytes),
	}
	return h.producer.SendMessage(&msg)
}
