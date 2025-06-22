package serverless_runtime

import (
	"net/http"

	"github.com/ArcherChu123/plugin-daemon/internal/core/plugin_manager/basic_runtime"
	"github.com/ArcherChu123/plugin-daemon/internal/utils/mapping"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities"
	"github.com/ArcherChu123/plugin-daemon/pkg/entities/plugin_entities"
)

type ServerlessPluginRuntime struct {
	basic_runtime.BasicChecksum
	plugin_entities.PluginRuntime

	// access url for the lambda function
	LambdaURL  string
	LambdaName string

	// listeners mapping session id to the listener
	listeners mapping.Map[string, *entities.Broadcast[plugin_entities.SessionMessage]]

	client *http.Client

	PluginMaxExecutionTimeout int // in seconds
}
