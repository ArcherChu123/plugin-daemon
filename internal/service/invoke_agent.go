package service

import (
	"github.com/ArcherChu123/plugin-daemon/internal/core/plugin_daemon"
	"github.com/ArcherChu123/plugin-daemon/internal/core/plugin_daemon/access_types"
	"github.com/ArcherChu123/plugin-daemon/internal/core/session_manager"
	"github.com/ArcherChu123/plugin-daemon/internal/utils/stream"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities/agent_entities"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities/plugin_entities"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities/requests"
	"github.com/gin-gonic/gin"
)

func InvokeAgentStrategy(
	r *plugin_entities.InvokePluginRequest[requests.RequestInvokeAgentStrategy],
	ctx *gin.Context,
	max_timeout_seconds int,
) {
	baseSSEWithSession(
		func(session *session_manager.Session) (*stream.Stream[agent_entities.AgentStrategyResponseChunk], error) {
			return plugin_daemon.InvokeAgentStrategy(session, &r.Data)
		},
		access_types.PLUGIN_ACCESS_TYPE_AGENT_STRATEGY,
		access_types.PLUGIN_ACCESS_ACTION_INVOKE_AGENT_STRATEGY,
		r,
		ctx,
		max_timeout_seconds,
	)
}
