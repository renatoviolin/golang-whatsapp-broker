package services

import (
	"broker/app/entity"
	"broker/infra/http_client"
	"broker/infra/kafka"
	"broker/infra/redis"
	"broker/util"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var CLIENT_TEST = entity.NewAgentConfig("Agent Teste", "agent_test", "agent_TESTE-out", "agent_TESTE-in", "agent_TEST-error", "cg-test")
var redisCache *redis.RedisClient
var producer *kafka.Producer
var httpClient *http_client.Client
var sender *Sender
var router RouterWebhook

func init() {
	util.LoadVars()
	redisCache = redis.NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	producer = kafka.NewProducer([]string{"localhost:9093"})
	httpClient = http_client.NewClient()
	sender = NewSendService(httpClient)
	router = NewRouterWebhook(producer, redisCache, sender)
	router.RegisterAgent(CLIENT_TEST)
}

// Route a Status message (do not update redis and send to ROUTER topic)
func Test_Route_RouterStatusMessage(t *testing.T) {
	payload := util.GetPayloadStatus1()
	redisCache.ClearAll()
	err := router.Dispatch(payload)
	require.NoError(t, err)
	_, err = redisCache.Get("5516993259256")
	require.Error(t, err)
}

// Route first message from agent to broker
// - Message must be set to ROUTER_ID
// - Menu must be sent to whatsapp
func Test_Route_InitialMessage(t *testing.T) {
	payload := util.GetPayload1()
	redisCache.ClearAll()
	err := router.Dispatch(payload)
	require.NoError(t, err)
	agentID, err := redisCache.Get("5516993259256")
	require.NoError(t, err)
	require.Equal(t, ROUTER_ID, agentID)
}

// Route response message with invalid agentID
// - Message must be kept to ROUTER_ID
// - Menu must be sent to whatsapp again
func Test_Route_RouterMessage_ChooseInvalidAgentID(t *testing.T) {
	payload := util.GetPayloadInvalidAgentID()
	err := router.Dispatch(payload)
	require.NoError(t, err)
	agentID, err := redisCache.Get("5516993259256")
	require.NoError(t, err)
	require.Equal(t, "router", agentID)
}

// Route message from agent to ROUTER_ID
// - Message must be set to CLIENT_A
// - Must send the welcome message from CLIENT_A
func Test_Route_RouterMessage_2(t *testing.T) {
	payload := util.GetPayload2()
	err := router.Dispatch(payload)
	require.NoError(t, err)
	agentID, err := redisCache.Get("5516993259256")
	require.NoError(t, err)
	require.Equal(t, CLIENT_TEST.ID, agentID)
}

// Route message from agent to CLIENT_A
// - Message must be published in CLIENT_A
func Test_Route_RouterMessage_2a(t *testing.T) {
	payload := util.GetPayload3()
	err := router.Dispatch(payload)
	require.NoError(t, err)
	agentID, err := redisCache.Get("5516993259256")
	require.NoError(t, err)
	require.Equal(t, CLIENT_TEST.ID, agentID)
}

// Route seconds message from agent to broker
// - Message already routed to to CLIENT_A
// - Message must be published in CLIENT_A_IN
// - Must update the Expires of the waID
func Test_Route_RouterMessage_3(t *testing.T) {
	payload := util.GetPayload3()
	err := router.Dispatch(payload)
	require.NoError(t, err)
	agentID, err := redisCache.Get("5516993259256")
	require.NoError(t, err)
	require.Equal(t, CLIENT_TEST.ID, agentID)
}

func Test_Route_RouterMessage_4_updateTTL(t *testing.T) {
	time.Sleep(3 * time.Second)
	payload := util.GetPayload3()

	ttlPrev, err := redisCache.GetTTL("5516993259256")
	require.NoError(t, err)
	agentID, err := redisCache.Get("5516993259256")
	require.NoError(t, err)
	require.Equal(t, CLIENT_TEST.ID, agentID)

	err = router.Dispatch(payload)
	require.NoError(t, err)
	agentID2, err := redisCache.Get("5516993259256")
	require.NoError(t, err)
	require.Equal(t, CLIENT_TEST.ID, agentID2)
	ttlCurrent, err := redisCache.GetTTL("5516993259256")
	require.NoError(t, err)
	require.Greater(t, ttlCurrent, ttlPrev)
}

// Route Sair message from agent to broker
// - Message must be set to router
func Test_Route_RouterMessage_Sair(t *testing.T) {
	payload := util.GetPayloadSair()
	err := router.Dispatch(payload)
	require.NoError(t, err)
	agentID, err := redisCache.Get("5516993259256")
	require.NoError(t, err)
	require.Equal(t, "router", agentID)
}
