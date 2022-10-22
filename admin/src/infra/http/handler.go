package http

import (
	"strconv"
	"whatsapp-admin/dto"

	"github.com/gin-gonic/gin"
)

func (h *HttpServer) health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func (h *HttpServer) createAgent(c *gin.Context) {
	var inputPayload dto.CreateAgentInput
	if err := c.ShouldBindJSON(&inputPayload); err != nil {
		e := newError("create-01", err.Error())
		c.AbortWithStatusJSON(400, e)
		return
	}
	id, err := h.agentController.Create(inputPayload)
	if err != nil {
		e := newError("create-02", err.Error())
		c.AbortWithStatusJSON(500, e)
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (h *HttpServer) updateAgent(c *gin.Context) {
	var inputPayload dto.UpdateAgentInput
	if err := c.ShouldBindJSON(&inputPayload); err != nil {
		e := newError("update-01", err.Error())
		c.AbortWithStatusJSON(400, e)
		return
	}
	err := h.agentController.Update(inputPayload)
	if err != nil {
		e := newError("update-02", err.Error())
		c.AbortWithStatusJSON(500, e)
		return
	}
	c.JSON(200, gin.H{"status": "success"})
}

func (h *HttpServer) deleteAgent(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		e := newError("delete-01", err.Error())
		c.AbortWithStatusJSON(400, e)
		return
	}

	err = h.agentController.Delete(intID)
	if err != nil {
		e := newError("delete-02", err.Error())
		c.AbortWithStatusJSON(500, e)
		return
	}
	c.JSON(200, gin.H{"status": "success"})
}

func (h *HttpServer) findAgentById(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		e := newError("find-01", err.Error())
		c.AbortWithStatusJSON(400, e)
		return
	}

	output, err := h.agentController.FindById(intID)
	if err != nil {
		e := newError("find-02", err.Error())
		c.AbortWithStatusJSON(500, e)
		return
	}
	c.JSON(200, gin.H{"status": "success", "data": output})
}

func (h *HttpServer) findAllAgent(c *gin.Context) {
	output, err := h.agentController.FindAll()
	if err != nil {
		e := newError("findAll-01", err.Error())
		c.AbortWithStatusJSON(500, e)
		return
	}
	c.JSON(200, gin.H{"status": "success", "data": output})
}

func (h *HttpServer) findAgentByName(c *gin.Context) {
	name := c.Param("name")
	output, err := h.agentController.FindByName(name)
	if err != nil {
		e := newError("findByName-01", err.Error())
		c.AbortWithStatusJSON(500, e)
		return
	}
	c.JSON(200, gin.H{"status": "success", "data": output})
}
