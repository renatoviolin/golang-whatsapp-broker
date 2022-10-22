package services

import (
	"encoding/json"
	"strings"
	"whatsapp-client/app/agent"
	"whatsapp-client/dto"
	"whatsapp-client/infra/custom_errors"
	"whatsapp-client/infra/kafka"
)

type ProducerService struct {
	kafkaProducer kafka.Producer
	agent         agent.AgentConfig
}

func NewProducerService(kafkaProducer kafka.Producer, agent agent.AgentConfig) ProducerService {
	return ProducerService{kafkaProducer: kafkaProducer, agent: agent}
}

func (h *ProducerService) SendToKafka(input dto.KafkaMessageOutput) error {
	if input.WaID == "" {
		return custom_errors.New("send-kafka", "empty waID")
	}
	if input.MessageType != "text" && input.MessageType != "list" {
		return custom_errors.New("send-kafka", "invalid message type")
	}
	if input.MessageType == "text" && input.Text == "" {
		return custom_errors.New("send-kafka", "empty text")
	}
	if input.MessageType == "list" {
		if input.ListBody == "" || input.ListTitle == "" || len(input.List) == 0 {
			return custom_errors.New("send-kafka", "invalid list parameters")
		}
		for _, v := range input.List {
			if strings.TrimSpace(v.ID) == "" || strings.TrimSpace(v.Title) == "" {
				return custom_errors.New("send-kafka", "list_item with empty ID or Title")
			}
		}
	}
	outputBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}
	_, _, err = h.kafkaProducer.Produce(outputBytes, h.agent.WriteTopic)
	if err != nil {
		return err
	}
	return nil
}
