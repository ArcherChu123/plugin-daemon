package service

import (
	"github.com/ArcherChu123/plugin-daemon/internal/core/plugin_daemon/backwards_invocation/transaction"
	"github.com/gin-gonic/gin"
)

func HandleAWSPluginTransaction(handler *transaction.AWSTransactionHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get session id from the context
		sessionId := c.Request.Header.Get("Dify-Plugin-Session-ID")

		handler.Handle(c, sessionId)
	}
}
