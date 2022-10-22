package http

import (
	"broker/controllers"
	"broker/infra/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	gin               *gin.Engine
	webhookController controllers.WebhookController
}

func NewHttpServer(webhookController controllers.WebhookController) HttpServer {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin := gin.Default()
	return HttpServer{
		gin:               gin,
		webhookController: webhookController,
	}
}

func (h *HttpServer) StartServer(port string) {
	// if err := h.gin.RunTLS(fmt.Sprintf("0.0.0.0:%s", port), "./certs/fullchain17.pem", "./certs/privkey17.pem"); err != nil {
	if err := h.gin.Run(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		logger.Fatal("router", "StartServer", fmt.Sprintf("%v", err))
	}
}

func (h *HttpServer) SetupRouter() {
	h.gin.GET("/", h.health)
	h.gin.POST("/webhook", h.receiveMessage)
	h.gin.GET("/webhook", h.verifyToken)
}
