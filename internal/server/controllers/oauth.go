package controllers

import (
	"time"

	"github.com/ArcherChu123/plugin-daemon/internal/service"
	"github.com/ArcherChu123/plugin-daemon/internal/types/app"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities/plugin_entities"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities/requests"
	"github.com/gin-gonic/gin"
)

func OAuthGetAuthorizationURL(config *app.Config) gin.HandlerFunc {
	type request = plugin_entities.InvokePluginRequest[requests.RequestOAuthGetAuthorizationURL]

	return func(c *gin.Context) {
		BindPluginDispatchRequest(
			c,
			func(ipr request) {
				service.OAuthGetAuthorizationURL(
					&ipr,
					c,
					time.Duration(config.PluginMaxExecutionTimeout)*time.Second,
				)
			},
		)
	}
}

func OAuthGetCredentials(config *app.Config) gin.HandlerFunc {
	type request = plugin_entities.InvokePluginRequest[requests.RequestOAuthGetCredentials]

	return func(c *gin.Context) {
		BindPluginDispatchRequest(c, func(ipr request) {
			service.OAuthGetCredentials(
				&ipr,
				c,
				time.Duration(config.PluginMaxExecutionTimeout)*time.Second,
			)
		})
	}
}
