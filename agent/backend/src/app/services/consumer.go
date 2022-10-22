package services

import (
	"encoding/json"
	"fmt"
	"time"
	"whatsapp-client/dto"
	"whatsapp-client/infra/custom_errors"
	"whatsapp-client/infra/logger"
	"whatsapp-client/infra/repository"

	"github.com/google/uuid"
)

type ConsumerService struct {
	repository repository.MessageRepository
}

func NewConsumerService(repository repository.MessageRepository) ConsumerService {
	return ConsumerService{repository: repository}
}

func (h *ConsumerService) DecodeSuccessMessage(inputBytes []byte) (err error) {
	var kafkaMessageInput dto.KafkaMessageInput
	if err = json.Unmarshal(inputBytes, &kafkaMessageInput); err != nil {
		e := custom_errors.New("decode-success", err.Error())
		return e.Err()
	}

	if kafkaMessageInput.Type == "text" || kafkaMessageInput.Type == "list" || kafkaMessageInput.Type == "initial" {
		messageRepository := dto.MessageRepository{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			Type:      kafkaMessageInput.Type,
			Body:      kafkaMessageInput.Text,
			IsRead:    false,
			WaID:      kafkaMessageInput.WaID,
		}
		if err = h.repository.Create(messageRepository); err != nil {
			return err
		}
	}

	logger.Info("consumer", fmt.Sprintf("type: %s  text: %s", kafkaMessageInput.Type, kafkaMessageInput.Text))
	return nil
}

func (h *ConsumerService) DecodeErrorMessage(inputBytes []byte) (err error) {
	var kafkaErrorInput dto.KafkaErrorInput
	if err = json.Unmarshal(inputBytes, &kafkaErrorInput); err != nil {
		e := custom_errors.New("decode-success", err.Error())
		return e.Err()
	}

	messageRepository := dto.MessageRepository{
		Id:        uuid.NewString(),
		CreatedAt: time.Now(),
		Type:      "error",
		Body:      kafkaErrorInput.Text,
		IsRead:    false,
		WaID:      kafkaErrorInput.WaID,
	}

	if err = h.repository.Create(messageRepository); err != nil {
		return err
	}

	logger.Info("consumer-error", kafkaErrorInput.Text)
	return nil
}
