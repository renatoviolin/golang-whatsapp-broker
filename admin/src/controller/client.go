package controller

import (
	"whatsapp-admin/app/services"
	"whatsapp-admin/dto"
)

type AgentController struct {
	service services.AgentService
}

func NewAgentController(service services.AgentService) AgentController {
	return AgentController{service: service}
}

func (h *AgentController) Create(inputPayload dto.CreateAgentInput) (int, error) {
	input := dto.Agent{
		Name:       inputPayload.Name,
		ReadTopic:  inputPayload.ReadTopic,
		WriteTopic: inputPayload.WriteTopic,
		ErrorTopic: inputPayload.ErrorTopic,
	}
	return h.service.Create(input)
}

func (h *AgentController) Update(inputPayload dto.UpdateAgentInput) error {
	input := dto.Agent{
		ID:         inputPayload.ID,
		Name:       inputPayload.Name,
		ReadTopic:  inputPayload.ReadTopic,
		WriteTopic: inputPayload.WriteTopic,
		ErrorTopic: inputPayload.ErrorTopic,
	}
	return h.service.Update(input)
}

func (h *AgentController) Delete(inputPayload int) error {
	return h.service.Delete(inputPayload)
}

func (h *AgentController) FindByName(inputPayload string) ([]*dto.Agent, error) {
	return h.service.FindByName(inputPayload)
}

func (h *AgentController) FindById(inputPayload int) (dto.Agent, error) {
	return h.service.FindById(inputPayload)
}

func (h *AgentController) FindAll() ([]*dto.Agent, error) {
	return h.service.All()
}
