package service

import (
	"github.com/ArcherChu123/plugin-daemon/internal/core/plugin_daemon"
	"github.com/ArcherChu123/plugin-daemon/internal/core/plugin_daemon/access_types"
	"github.com/ArcherChu123/plugin-daemon/internal/core/session_manager"
	"github.com/ArcherChu123/plugin-daemon/internal/utils/stream"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities/plugin_entities"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities/requests"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities/tool_entities"
	"github.com/gin-gonic/gin"
)

func InvokeTool(
	r *plugin_entities.InvokePluginRequest[requests.RequestInvokeTool],
	ctx *gin.Context,
	max_timeout_seconds int,
) {
	baseSSEWithSession(
		func(session *session_manager.Session) (*stream.Stream[tool_entities.ToolResponseChunk], error) {
			return plugin_daemon.InvokeTool(session, &r.Data)
		},
		access_types.PLUGIN_ACCESS_TYPE_TOOL,
		access_types.PLUGIN_ACCESS_ACTION_INVOKE_TOOL,
		r,
		ctx,
		max_timeout_seconds,
	)
}
