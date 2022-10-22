package services

import (
	"broker/infra/http_client"
	"broker/infra/kafka"
	"broker/infra/redis"
	"broker/util"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var routeClient RouteAgentMessages

func init() {
	util.LoadVars()
	redisCache = redis.NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	redisCache.ClearAll()
	producer = kafka.NewProducer([]string{"localhost:9093"})
	httpClient = http_client.NewClient()
	routeClient = NewRouteAgentMessages(producer, httpClient, redisCache, CLIENT_TEST)
}

func Test_Dispatch_TextMessage_Valid(t *testing.T) {
	redisCache.Save("5516993259256", "agent_test")
	inputBytes := []byte(`
		{
			"waID": "5516993259256",
			"message_type": "text",
			"text": "mensagem de teste do dispatch",
			"list_title": "",
			"list_items": []
		}
	`)
	err := routeClient.Dispatch(inputBytes)
	require.NoError(t, err)
	redisCache.ClearAll()
}

func Test_Dispatch_TextMessage_Invalid(t *testing.T) {
	redisCache.Save("5516993259256", "agent_test")
	inputBytes := []byte(`
		{
			"waID": "",
			"message_type": "text",
			"text": "mensagem de teste do dispatch",
			"list_title": "",
			"list_items": []
		}
	`)
	err := routeClient.Dispatch(inputBytes)
	require.Error(t, err)

	inputBytes = []byte(`
		{
			"waID": "5516993259256",
			"message_type": "text",
			"text": "",
			"list_title": "",
			"list_items": []
		}
	`)
	err = routeClient.Dispatch(inputBytes)
	require.Error(t, err)
	redisCache.ClearAll()
}

func Test_Dispatch_ListMessage_Valid(t *testing.T) {
	redisCache.Save("5516993259256", "agent_test")
	inputBytes := []byte(`
		{
			"waID": "5516993259256",
			"message_type": "list",
			"text": "",
			"list_title": "Titulo do Header",
			"list_body": "Conteúdo Header",
			"list_items": [
				{
					"id": "list_1",
					"title": "Opção 1"
				},
				{
					"id": "list_2",
					"title": "Opção 2"
				}
			]
		}
	`)
	err := routeClient.Dispatch(inputBytes)
	require.NoError(t, err)
	redisCache.ClearAll()
}

func Test_Dispatch_ListMessage_Invalid(t *testing.T) {
	redisCache.Save("5516993259256", "agent_test")
	inputBytes := []byte(`
		{
			"waID": "5516993259256",
			"message_type": "list",
			"text": "",
			"list_title": "Titulo do Header",
			"list_body": "Conteúdo Header",
			"list_items": [
				{
					"id": "",
					"title": "Opção 1"
				},
				{
					"id": "list_2",
					"title": "Opção 2"
				}
			]
		}
	`)
	err := routeClient.Dispatch(inputBytes)
	require.Error(t, err)

	inputBytes = []byte(`
		{
			"waID": "5516993259256",
			"message_type": "list",
			"text": "",
			"list_title": "Titulo do Header",
			"list_body": "Conteúdo Header",
			"list_items": [
				{
					"id": "list_1",
					"title": "Opção 1"
				},
				{
					"id": "list_1",
					"title": "Opção 2"
				}
			]
		}
	`)
	err = routeClient.Dispatch(inputBytes)
	require.Error(t, err)
	redisCache.ClearAll()
}
