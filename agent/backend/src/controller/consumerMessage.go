package controller

import (
	"whatsapp-client/app/services"
)

type ConsumerController struct {
	consumerService services.ConsumerService
}

func NewConsumerController(consumerService services.ConsumerService) ConsumerController {
	return ConsumerController{consumerService: consumerService}
}

func (h *ConsumerController) HandleMessage(input []byte) (err error) {
	return h.consumerService.DecodeSuccessMessage(input)
}
