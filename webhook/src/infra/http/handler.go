package http

import (
	"broker/infra/logger"
	"errors"
	"io"

	"github.com/gin-gonic/gin"
)

func (h *HttpServer) health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func (h *HttpServer) receiveMessage(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("handler", "receive-message", err.Error())
		c.JSON(500, err.Error())
		return
	}

	err = h.webhookController.ReceiveFromWhatsapp(jsonData)
	if err != nil {
		logger.Error("handler", "receive-from-whatsapp", err.Error())
		c.JSON(500, err.Error())
		return
	}
	c.Writer.Write([]byte("ok"))
}

func (h *HttpServer) verifyToken(c *gin.Context) {
	mode := c.Query("hub.mode")
	challenge := c.Query("hub.challenge")
	token := c.Query("hub.verify_token")

	if mode == "" || challenge == "" || token == "" || mode != "subscribe" {
		err := errors.New("invalid parameters")
		logger.Error("handlers", "receiveMessage", err.Error())
		c.JSON(400, err.Error())
		return
	}

	err := h.webhookController.VerifyToken(token)
	if err != nil {
		err := errors.New("invalid token")
		logger.Error("handlers", "token", err.Error())
		c.JSON(400, err.Error())
		return
	}
	c.Writer.Write([]byte(challenge))
}
