package controllers

import (
	"github.com/ArcherChu123/plugin-daemon/internal/manifest"
	"github.com/ArcherChu123/plugin-daemon/internal/types/app"
	"github.com/ArcherChu123/plugin-daemon/internal/utils/routine"
	"github.com/gin-gonic/gin"
)

func HealthCheck(app *app.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":      "ok",
			"pool_status": routine.FetchRoutineStatus(),
			"version":     manifest.VersionX,
			"build_time":  manifest.BuildTimeX,
			"platform":    app.Platform,
		})
	}
}
