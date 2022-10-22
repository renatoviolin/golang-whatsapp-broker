package main

import (
	"os"
	"whatsapp-client/app/agent"
	"whatsapp-client/app/services"
	"whatsapp-client/controller"
	"whatsapp-client/infra/http"
	"whatsapp-client/infra/kafka"
	"whatsapp-client/infra/logger"
	"whatsapp-client/infra/repository"
	"whatsapp-client/util"

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
	agentConfig := agent.NewAgentConfig(os.Getenv("READ_TOPIC"), os.Getenv("WRITE_TOPIC"), os.Getenv("ERROR_TOPIC"), "agent_A-cg")
	mongoClient, err := repository.NewMongoClient(os.Getenv("MONGO_CONNECTION"), os.Getenv("MONGO_DATABASE"), 5)
	if err != nil {
		logger.Fatal("main", err.Error())
	}
	repository := repository.NewMessageRepository(mongoClient, os.Getenv("DIALOG_COLLECTION"))

	// Setup producer
	kafkaProducer := kafka.NewProducer([]string{"localhost:9093"})
	produceService := services.NewProducerService(*kafkaProducer, agentConfig)
	responseController := controller.NewProducerController(produceService)

	// Setup consumer to read messages from broker
	consumerService := services.NewConsumerService(repository)
	consumerController := controller.NewConsumerController(consumerService)
	consumerHandler := NewConsumerHandler(&consumerController)
	kafkaConsumer := kafka.NewConsumer([]string{"localhost:9093"}, agentConfig.ReadTopic, agentConfig.ConsumerGroup, consumerHandler)
	kafkaConsumer.StartConsumer()

	// Setup consumer to read errors from broker
	consumerErrorController := controller.NewConsumerErrorController(consumerService)
	consumerFromRepository := controller.NewConsumerFromRepositoryController(repository)
	consumerErrorHandler := NewConsumerHandler(&consumerErrorController)
	kafkaConsumerError := kafka.NewConsumer([]string{"localhost:9093"}, agentConfig.ErrorTopic, agentConfig.ConsumerGroup, consumerErrorHandler)
	kafkaConsumerError.StartConsumer()

	httpServer := http.NewHttpServer(responseController, consumerFromRepository)
	httpServer.SetupRoutes()
	logger.Info("main", "server running: "+os.Getenv("PORT"))
	httpServer.Listen(os.Getenv("PORT"))
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
			err := h.handler.HandleMessage(message.Value)
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
