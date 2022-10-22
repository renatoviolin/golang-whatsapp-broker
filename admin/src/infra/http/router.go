package http

import (
	"fmt"
	"whatsapp-admin/controller"
	"whatsapp-admin/infra/logger"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	gin             *gin.Engine
	agentController controller.AgentController
}

func NewHttpServer(agentController controller.AgentController) HttpServer {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	engine := gin.New()
	engine.Use(StructuredLogger(&logger.Log), gin.Recovery())

	return HttpServer{
		gin:             engine,
		agentController: agentController,
	}
}

func (h *HttpServer) Listen(port string) {
	if err := h.gin.Run(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		logger.Fatal("start-server", fmt.Sprintf("%v", err))
	}
}

func (h *HttpServer) SetupRoutes() {
	h.gin.GET("/", h.health)
	h.gin.POST("/api/v1/agent/create", h.createAgent)
	h.gin.POST("/api/v1/agent/update", h.updateAgent)
	h.gin.DELETE("/api/v1/agent/delete/:id", h.deleteAgent)
	h.gin.GET("/api/v1/agent", h.findAllAgent)
	h.gin.GET("/api/v1/agent/:id", h.findAgentById)
	h.gin.GET("/api/v1/agent/by-name/:name", h.findAgentByName)
}
