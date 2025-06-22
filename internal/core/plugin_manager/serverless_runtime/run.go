package serverless_runtime

import "github.com/ArcherChu123/plugin-daemon/pkg/entities/plugin_entities"

func (r *ServerlessPluginRuntime) StartPlugin() error {
	return nil
}

func (r *ServerlessPluginRuntime) Wait() (<-chan bool, error) {
	return nil, nil
}

func (r *ServerlessPluginRuntime) Type() plugin_entities.PluginRuntimeType {
	return plugin_entities.PLUGIN_RUNTIME_TYPE_SERVERLESS
}
