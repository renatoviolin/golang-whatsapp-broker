package controller

import (
	"whatsapp-client/app/services"
)

type ConsumerErrorController struct {
	consumerService services.ConsumerService
}

func NewConsumerErrorController(consumerService services.ConsumerService) ConsumerErrorController {
	return ConsumerErrorController{consumerService: consumerService}
}

func (h *ConsumerErrorController) HandleMessage(input []byte) error {
	return h.consumerService.DecodeErrorMessage(input)

}
