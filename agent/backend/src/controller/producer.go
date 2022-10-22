package controller

import (
	"whatsapp-client/app/services"
	"whatsapp-client/dto"
)

type ProducerController struct {
	producerService services.ProducerService
}

func NewProducerController(producerService services.ProducerService) ProducerController {
	return ProducerController{producerService: producerService}
}

func (h *ProducerController) Produce(input dto.HttpMessageInputPayload) error {
	var listItems []dto.ListItem
	for _, v := range input.ListItems {
		listItems = append(listItems, dto.ListItem{ID: v.ID, Title: v.Title})
	}
	kafkaMessageOutput := dto.KafkaMessageOutput{
		WaID:        input.WaID,
		MessageType: input.Type,
		Text:        input.Text,
		ListTitle:   "List Title",
		ListBody:    "List Body",
		List:        listItems,
	}
	return h.producerService.SendToKafka(kafkaMessageOutput)
}
