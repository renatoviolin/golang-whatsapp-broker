package main

import (
	"broker/app/services"
	"broker/controllers"
	"broker/infra/http"
	"broker/infra/kafka"
	"broker/infra/logger"
	"broker/util"
	"os"
)

func init() {
	logger.InitLog()
	err := util.LoadVars()
	if err != nil {
		logger.Error("main", "load-vars", err.Error())
	}
}

func main() {
	servers := []string{os.Getenv("KAFKA_SERVER")}
	topic := os.Getenv("KAFKA_TOPIC")
	kafkaProducer := kafka.NewProducer(servers, topic)
	receiveService := services.NewReceiveService(kafkaProducer)
	tokenService := services.NewTokenService()
	webhookController := controllers.NewWebhookController(receiveService, tokenService)

	server := http.NewHttpServer(webhookController)
	server.SetupRouter()
	server.StartServer(os.Getenv("PORT"))
}
