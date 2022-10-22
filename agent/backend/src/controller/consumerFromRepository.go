package controller

import (
	"whatsapp-client/dto"
	"whatsapp-client/infra/repository"
)

type ConsumerFromRepositoryController struct {
	repository repository.MessageRepository
}

func NewConsumerFromRepositoryController(repository repository.MessageRepository) ConsumerFromRepositoryController {
	return ConsumerFromRepositoryController{repository: repository}
}

func (h *ConsumerFromRepositoryController) ReadMessages(waID string) (output []dto.MessageRepository, err error) {
	return h.repository.FindAll(waID)
}

func (h *ConsumerFromRepositoryController) MarkAsRead(id string) error {
	return h.repository.MarkAsRead(id)
}
