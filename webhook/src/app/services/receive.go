package services

import (
	"broker/infra/kafka"
	"broker/infra/logger"
)

type ReceiveService struct {
	producer *kafka.Producer
}

func NewReceiveService(producer *kafka.Producer) ReceiveService {
	return ReceiveService{
		producer: producer,
	}
}

func (h *ReceiveService) Receive(input []byte) error {
	_, _, err := h.producer.Produce(input)
	if err == nil {
		logger.Info("services", "receive", "message receveid with success")
	}
	return err
}
