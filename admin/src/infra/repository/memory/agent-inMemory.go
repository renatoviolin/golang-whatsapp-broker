package memory

import (
	"errors"
	"math/rand"
	"whatsapp-admin/dto"
)

type AgentMemoryRepository struct {
	database []dto.Agent
}

func NewAgentMemoryRepository() *AgentMemoryRepository {
	return &AgentMemoryRepository{}
}

func (h *AgentMemoryRepository) FindById(id int) (dto.Agent, error) {
	for _, el := range h.database {
		if el.ID == id {
			return el, nil
		}
	}
	return dto.Agent{}, errors.New("resource not found")
}

func (h *AgentMemoryRepository) Create(agent dto.Agent) (int, error) {
	id := rand.Intn(1000000000)
	agent.ID = id
	h.database = append(h.database, agent)
	return id, nil
}

func (h *AgentMemoryRepository) Update(agent dto.Agent) error {
	for i, el := range h.database {
		if el.ID == agent.ID {
			h.database[i] = agent
			return nil
		}
	}
	return errors.New("resource not found")
}

func (h *AgentMemoryRepository) Delete(id int) error {
	for i, el := range h.database {
		if el.ID == id {
			h.database = append(h.database[:i], h.database[i+1:]...)
			return nil
		}
	}
	return errors.New("resource not found")
}

func (h *AgentMemoryRepository) FindByName(name string) ([]*dto.Agent, error) {
	var result []*dto.Agent
	for i := range h.database {
		if h.database[i].Name == name {
			result = append(result, &h.database[i])
		}
	}
	if len(result) == 0 {
		return result, errors.New("resource not found")
	}
	return result, nil
}

func (h *AgentMemoryRepository) All() ([]*dto.Agent, error) {
	var result []*dto.Agent
	for i := range h.database {
		result = append(result, &h.database[i])
	}
	if len(result) == 0 {
		return result, errors.New("resource not found")
	}
	return result, nil
}
