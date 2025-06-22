package controllers

import (
	"github.com/ArcherChu123/plugin-daemon/internal/service"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities/requests"
	"github.com/gin-gonic/gin"
)

func GetRemoteDebuggingKey(c *gin.Context) {
	BindRequest(
		c, func(request requests.RequestGetRemoteDebuggingKey) {
			c.JSON(200, service.GetRemoteDebuggingKey(request.TenantID))
		},
	)
}
