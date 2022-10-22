package memory

import (
	"testing"
	"whatsapp-admin/dto"
	"whatsapp-admin/util"

	"github.com/stretchr/testify/require"
)

var repository = NewAgentMemoryRepository()

func init() {
	util.LoadVars()
}

func Test_Insert(t *testing.T) {
	agent := dto.Agent{
		Name:       "nome-1",
		ReadTopic:  "read-topic-1",
		WriteTopic: "write-topic-1",
		ErrorTopic: "error-topic-1",
	}

	id, err := repository.Create(agent)
	require.NoError(t, err)
	require.Greater(t, id, 0)
}

func Test_GetandUpdate(t *testing.T) {
	agent := dto.Agent{
		Name:       "nome-1",
		ReadTopic:  "read-topic-1",
		WriteTopic: "write-topic-1",
		ErrorTopic: "error-topic-1",
	}
	id, err := repository.Create(agent)
	require.NoError(t, err)

	// get agent
	agentDB, err := repository.FindById(id)
	require.NoError(t, err)

	// update agent
	agentDB.Name = "nome-atualizado-1"
	agentDB.ReadTopic = "read_topic-atualizado-1"
	agentDB.WriteTopic = "write_topic-atualizado-1"
	agentDB.ErrorTopic = "error_topic-atualizado-1"
	err = repository.Update(agentDB)
	require.NoError(t, err)

	// get agent again
	agentDB2, err := repository.FindById(id)
	require.NoError(t, err)
	require.Equal(t, "nome-atualizado-1", agentDB2.Name)

	// get Invalid id
	_, err = repository.FindById(-1)
	require.Error(t, err)
}

func Test_FindByName(t *testing.T) {
	// create agent
	agent := dto.Agent{
		Name:       "nome-para-teste",
		ReadTopic:  "read-topic-1",
		WriteTopic: "write-topic-1",
		ErrorTopic: "error-topic-1",
	}
	_, err := repository.Create(agent)
	require.NoError(t, err)

	// get agent
	agentDB, err := repository.FindByName("nome-para-teste")
	require.NoError(t, err)
	require.Greater(t, len(agentDB), 0)

	// get Invalid name
	_, err = repository.FindByName("xxxxx")
	require.Error(t, err)
}

func Test_GetAll(t *testing.T) {
	agentDB, err := repository.All()
	require.NoError(t, err)
	require.Greater(t, len(agentDB), 0)
}

func Test_DeleteAll(t *testing.T) {
	// create agent
	agent := dto.Agent{
		Name:       "nome-1",
		ReadTopic:  "read-topic-1",
		WriteTopic: "write-topic-1",
		ErrorTopic: "error-topic-1",
	}
	id, err := repository.Create(agent)
	require.NoError(t, err)

	// delete agent
	err = repository.Delete(id)
	require.NoError(t, err)

	_, err = repository.FindById(id)
	require.Error(t, err)

	// delete invalid id
	err = repository.Delete(-1)
	require.Error(t, err)
}
