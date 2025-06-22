package controllers

import (
	"net/http"

	"github.com/ArcherChu123/plugin-daemon/internal/service"
	"github.com/gin-gonic/gin"
)

func ListModels(c *gin.Context) {
	BindRequest(c, func(request struct {
		TenantID string `uri:"tenant_id" validate:"required"`
		Page     int    `form:"page" validate:"required,min=1"`
		PageSize int    `form:"page_size" validate:"required,min=1,max=256"`
	}) {
		c.JSON(http.StatusOK, service.ListModels(request.TenantID, request.Page, request.PageSize))
	})
}
