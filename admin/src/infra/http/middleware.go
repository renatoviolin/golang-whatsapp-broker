package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func StructuredLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		var logEvent *zerolog.Event
		if c.Writer.Status() >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}

		logEvent.Str("from", "http").
			Str("method", param.Method).
			Str("path", param.Path).
			Int("status", param.StatusCode).
			Msg(param.ErrorMessage)
	}
}
