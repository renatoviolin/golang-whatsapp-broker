package http

import (
	"encoding/json"
	"net/http"
	"os"
	"whatsapp-client/dto"
	"whatsapp-client/infra/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *HttpServer) wsHandleMessage(c *gin.Context) {
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("ws-upgrader", newError("ws-upgrader", err.Error()).Message)
		return
	}
	defer wsConn.Close()

	for {
		_, p, err := wsConn.ReadMessage()
		if err != nil {
			logger.Error("ws-ReadMessage", newError("ws-ReadMessage", err.Error()).Message)
			break
		}

		var inputPayload dto.HttpMessageInputPayload
		inputPayload.WaID = os.Getenv("DEMO_WAID")
		if err = json.Unmarshal(p, &inputPayload); err != nil {
			logger.Error("ws-unmarshal", newError("ws-unmarshal", err.Error()).Message)
			continue
		}

		switch inputPayload.WSAction {
		case "send":
			err = h.producerController.Produce(inputPayload)
			if err != nil {
				logger.Error("ws-producer", newError("ws-producer", err.Error()).Message)
			}

		case "read":
			messagesNotRead, err := h.consumerFromRepositoryController.ReadMessages(inputPayload.WaID)
			if err != nil {
				logger.Error("from-repository", newError("from-repository", err.Error()).Message)
				sendWebsocketMessage([]byte(err.Error()), wsConn)
				break
			}
			for _, v := range messagesNotRead {
				bytes, _ := json.Marshal(v)
				err := sendWebsocketMessage(bytes, wsConn)
				if err == nil {
					err := h.consumerFromRepositoryController.MarkAsRead(v.Id)
					if err != nil {
						logger.Error("mark-as-read", err.Error())
					}
				}
			}
		}
	}
}

func sendWebsocketMessage(message []byte, ws *websocket.Conn) error {
	if err := ws.WriteMessage(1, message); err != nil {
		logger.Error("ws-writeMessage", newError("ws-writeMessage", err.Error()).Message)
		return err
	}
	return nil
}
