package repository

import "whatsapp-admin/dto"

type AgentRepositoy interface {
	Create(dto.Agent) (int, error)
	Update(dto.Agent) error
	Delete(int) error
	FindById(int) (dto.Agent, error)
	FindByName(string) ([]*dto.Agent, error)
	All() ([]*dto.Agent, error)
}
