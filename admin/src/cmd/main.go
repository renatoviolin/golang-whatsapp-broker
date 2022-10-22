package main

import (
	"os"
	"whatsapp-admin/app/services"
	"whatsapp-admin/controller"
	"whatsapp-admin/infra/http"
	"whatsapp-admin/infra/logger"
	"whatsapp-admin/infra/repository/postgres"
	"whatsapp-admin/util"
)

func init() {
	logger.InitLog()
	err := util.LoadVars()
	if err != nil {
		logger.Error("load-vars", err.Error())
	}
}

func main() {
	// Database
	repository, err := postgres.NewAgentRepository(os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatal("main", err.Error())
	}

	// Service
	agentService := services.NewAgenteService(repository)

	// Controller
	agentController := controller.NewAgentController(agentService)

	// HTTP
	httpServer := http.NewHttpServer(agentController)
	httpServer.SetupRoutes()
	logger.Info("main", "server running: "+os.Getenv("PORT"))
	httpServer.Listen(os.Getenv("PORT"))
}
