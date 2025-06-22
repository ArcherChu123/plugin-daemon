package agent_entities

import "github.com/ArcherChu123/plugin-daemon/pkg/entities/tool_entities"

type AgentStrategyResponseChunk struct {
	tool_entities.ToolResponseChunk `json:",inline"`
}
