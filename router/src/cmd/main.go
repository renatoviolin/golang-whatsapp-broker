package main

import (
	"broker/app/entity"
	"broker/app/services"
	"broker/controller"
	"broker/infra/http_client"
	"broker/infra/kafka"
	"broker/infra/logger"
	"broker/infra/redis"
	"broker/infra/repository/postgres"
	"broker/util"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

func init() {
	logger.InitLog()
	err := util.LoadVars()
	if err != nil {
		logger.Error("load-vars", err.Error())
	}
}

func main() {
	// infra
	servers := []string{os.Getenv("KAFKA_SERVER")}
	producer := kafka.NewProducer(servers)
	redis := redis.NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	httpagent := http_client.NewClient()

	// services
	sender := services.NewSendService(httpagent)
	routerService := services.NewRouterWebhook(producer, redis, sender)

	// agents
	agentRepository, err := postgres.NewAgentRepository(os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatal("main", err.Error())
	}

	agents, err := agentRepository.All()
	if err != nil {
		logger.Fatal("main", err.Error())
	}

	for _, v := range agents {
		ag := entity.NewAgentConfig(v.Name, v.ReadTopic, v.ReadTopic, v.WriteTopic, v.ErrorTopic, fmt.Sprintf("cg_%d", v.ID))
		routerService.RegisterAgent(ag)
		logger.Info("agents", fmt.Sprintf("registering %s", ag.ID))

		routeragent := services.NewRouteAgentMessages(producer, httpagent, redis, ag)
		controlleragent := controller.NewAgentController(routeragent)
		kafkaagentHandler := NewConsumerHandler(&controlleragent)
		kafkaagentConsumer := kafka.NewConsumer(servers, ag.WriteTopic, ag.ConsumerGroup, kafkaagentHandler)
		kafkaagentConsumer.StartConsumer()
	}

	// router
	controllerRouter := controller.NewRouterController(&routerService)
	kafkaRouterHandler := NewConsumerHandler(controllerRouter)
	kafkaRouterConsumer := kafka.NewConsumer(servers, os.Getenv("KAFKA_WEBHOOK_TOPIC"), "cg-router", kafkaRouterHandler)
	kafkaRouterConsumer.StartConsumer()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
	logger.Info("main", "terminating via signal")
}

type ConsumerHandler struct {
	handler controller.IController
}

func NewConsumerHandler(handler controller.IController) *ConsumerHandler {
	return &ConsumerHandler{handler: handler}
}

func (h *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			err := h.handler.HandleMessageFromKafka(message.Value)
			if err == nil {
				session.MarkMessage(message, "")
			} else {
				logger.Error("consume-claim", err.Error())
			}
		case <-session.Context().Done():
			return nil
		}
	}
}

func (h *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
