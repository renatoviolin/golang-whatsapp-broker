package services

// import (
// 	"broker/app/entity"
// 	"broker/infra/http_agent"
// 	"broker/infra/redis"
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func Test_SendText(t *testing.T) {
// 	httpagent := http_agent.Newagent()
// 	redis := redis.NewRedisagent(os.Getenv("REDIS_CONNECTION"))
// 	service := NewSendService(httpagent, redis)
// 	redis.Save("5516993259256", "agent_test")
// 	err := service.SendText("enviando a partir do golang", "5516993259256", "agent_test")
// 	require.NoError(t, err)
// }

// func Test_SendText_Differentagent(t *testing.T) {
// 	httpagent := http_agent.Newagent()
// 	redis := redis.NewRedisagent(os.Getenv("REDIS_CONNECTION"))
// 	service := NewSendService(httpagent, redis)
// 	redis.Save("5516993259256", "router")
// 	err := service.SendText("enviando a partir do golang", "5516993259256", "agent_test")
// 	require.Error(t, err)
// }

// func Test_InvalidNumber(t *testing.T) {
// 	httpagent := http_agent.Newagent()
// 	redis := redis.NewRedisagent(os.Getenv("REDIS_CONNECTION"))
// 	service := NewSendService(httpagent, redis)
// 	err := service.SendText("enviando a partir do golang", "5516993259255", "agent_test")
// 	require.Error(t, err)
// }

// func Test_SendList(t *testing.T) {
// 	httpagent := http_agent.Newagent()
// 	redis := redis.NewRedisagent(os.Getenv("REDIS_CONNECTION"))
// 	service := NewSendService(httpagent, redis)
// 	redis.Save("5516993259256", "agent_test")

// 	actions := []entity.Row{
// 		{ID: "agent_a", Title: "agente A"},
// 		{ID: "agent_b", Title: "agente B"},
// 		{ID: "agent_c", Title: "agente C"},
// 	}

// 	err := service.SendList("Escolha o Sistema", actions, "5516993259256", "agent_test")
// 	require.NoError(t, err)
// }

// func Test_agentNotAssociated(t *testing.T) {
// 	httpagent := http_agent.Newagent()
// 	redis := redis.NewRedisagent(os.Getenv("REDIS_CONNECTION"))
// 	service := NewSendService(httpagent, redis)
// 	err := service.SendText("enviando a partir do golang", "5516993259255", "agent_test")
// 	require.Error(t, err)
// }
