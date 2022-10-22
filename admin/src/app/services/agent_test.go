package services

import (
	"testing"
	"whatsapp-admin/dto"
	"whatsapp-admin/infra/repository"
	"whatsapp-admin/infra/repository/memory"
	"whatsapp-admin/util"

	"github.com/stretchr/testify/require"
)

var agentRepository repository.AgentRepositoy

func init() {
	util.LoadVars()
	// repository, _ = postgres.NewAgentRepository(os.Getenv("DATABASE_URL"))
	agentRepository = memory.NewAgentMemoryRepository()
}

func Test_Create_Valid(t *testing.T) {

	service := NewAgenteService(agentRepository)
	payload := dto.Agent{
		Name:       "nome-1",
		ReadTopic:  "read-topic-1",
		WriteTopic: "write-topic-1",
		ErrorTopic: "error-topic-1",
	}
	id, err := service.Create(payload)
	require.NoError(t, err)
	require.Greater(t, id, 0)
}

func Test_Create_Invalid(t *testing.T) {
	service := NewAgenteService(agentRepository)
	payload := dto.Agent{
		Name:       "",
		ReadTopic:  "read-topic-1",
		WriteTopic: "write-topic-1",
		ErrorTopic: "error-topic-1",
	}
	_, err := service.Create(payload)
	require.Error(t, err)
}

func Test_Update_Valid(t *testing.T) {
	service := NewAgenteService(agentRepository)
	payload := dto.Agent{
		Name:       "nome-1",
		ReadTopic:  "read-topic-1",
		WriteTopic: "write-topic-1",
		ErrorTopic: "error-topic-1",
	}
	id, err := service.Create(payload)
	require.NoError(t, err)
	payload.Name = "novo nome"
	payload.ID = id

	err = service.Update(payload)
	require.NoError(t, err)
}

func Test_Delete_Valid(t *testing.T) {
	service := NewAgenteService(agentRepository)
	payload := dto.Agent{
		Name:       "nome-1",
		ReadTopic:  "read-topic-1",
		WriteTopic: "write-topic-1",
		ErrorTopic: "error-topic-1",
	}
	id, err := service.Create(payload)
	require.NoError(t, err)

	err = service.Delete(id)
	require.NoError(t, err)

	err2 := service.Delete(id)
	require.Error(t, err2)
}

func Test_ById(t *testing.T) {
	service := NewAgenteService(agentRepository)
	payload := dto.Agent{
		Name:       "nome-1",
		ReadTopic:  "read-topic-1",
		WriteTopic: "write-topic-1",
		ErrorTopic: "error-topic-1",
	}
	id, err := service.Create(payload)
	require.NoError(t, err)
	payload.Name = "novo nome"
	payload.ID = id

	output, err := service.FindById(id)
	require.NoError(t, err)
	require.Equal(t, id, output.ID)
}

func Test_ByName(t *testing.T) {
	service := NewAgenteService(agentRepository)
	payload := dto.Agent{
		Name:       "nome-1",
		ReadTopic:  "read-topic-1",
		WriteTopic: "write-topic-1",
		ErrorTopic: "error-topic-1",
	}
	id, err := service.Create(payload)
	require.NoError(t, err)
	payload.Name = "novo nome"
	payload.ID = id

	output, err := service.FindByName("nome-1")
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(output), 1)
}

func Test_All(t *testing.T) {
	service := NewAgenteService(agentRepository)
	output, err := service.All()
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(output), 1)
}
