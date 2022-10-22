package services

import (
	"errors"
	"whatsapp-admin/dto"
	"whatsapp-admin/infra/repository"
)

type AgentService struct {
	repository repository.AgentRepositoy
}

func NewAgenteService(repository repository.AgentRepositoy) AgentService {
	return AgentService{repository: repository}
}

func (h *AgentService) Create(input dto.Agent) (id int, err error) {
	if err := isValid(input); err != nil {
		return id, err
	}
	id, err = h.repository.Create(input)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (h *AgentService) Update(input dto.Agent) (err error) {
	if err := isValid(input); err != nil {
		return err
	}
	if err = h.repository.Update(input); err != nil {
		return err
	}
	return nil
}

func (h *AgentService) Delete(input int) (err error) {
	if err = h.repository.Delete(input); err != nil {
		return err
	}
	return nil
}

func (h *AgentService) All() ([]*dto.Agent, error) {
	output, err := h.repository.All()
	if err != nil {
		return output, err
	}
	return output, nil
}

func (h *AgentService) FindById(id int) (dto.Agent, error) {
	output, err := h.repository.FindById(id)
	if err != nil {
		return output, err
	}
	return output, nil
}

func (h *AgentService) FindByName(name string) ([]*dto.Agent, error) {
	output, err := h.repository.FindByName(name)
	if err != nil {
		return output, err
	}
	return output, nil
}

func isValid(agent dto.Agent) error {
	if agent.Name == "" || agent.ReadTopic == "" || agent.WriteTopic == "" || agent.ErrorTopic == "" {
		return errors.New("invalid values")
	}
	return nil
}
