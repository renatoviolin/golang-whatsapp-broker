package http

import (
	"fmt"
	"whatsapp-client/controller"
	"whatsapp-client/infra/logger"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	gin                              *gin.Engine
	producerController               controller.ProducerController
	consumerFromRepositoryController controller.ConsumerFromRepositoryController
}

func NewHttpServer(producerController controller.ProducerController, consumerFromRepositoryController controller.ConsumerFromRepositoryController) HttpServer {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin := gin.Default()
	gin.Use(CORS)
	return HttpServer{
		gin:                              gin,
		producerController:               producerController,
		consumerFromRepositoryController: consumerFromRepositoryController,
	}
}

func (h *HttpServer) Listen(port string) {
	if err := h.gin.Run(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		logger.Fatal("start-server", fmt.Sprintf("%v", err))
	}
}

func (h *HttpServer) SetupRoutes() {
	h.gin.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	h.gin.GET("/websocket", h.wsHandleMessage)
}

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(200)
	}
}
