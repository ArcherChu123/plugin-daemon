package config

type AppConfig struct {
	Server struct {
		Port int    `env:"SERVER_PORT" default:"8080"`
		Env  string `env:"ENVIRONMENT" default:"development"`
	}

	PluginStorage struct {
		Type      string `env:"PLUGIN_STORAGE_TYPE" default:"local"` // local / oss
		LocalRoot string `env:"PLUGIN_STORAGE_LOCAL_ROOT"`
	}

	PythonRuntime struct {
		InterpreterPath string `env:"PYTHON_INTERPRETER_PATH" required:"true"`
		UvPath          string `env:"UV_PATH"`
		EnvInitTimeout  int    `env:"PYTHON_ENV_INIT_TIMEOUT" default:"60"`
	}

	Serverless struct {
		ConnectorURL     string `env:"PLUGIN_SERVERLESS_CONNECTOR_URL"`
		ConnectorAPIKey  string `env:"PLUGIN_SERVERLESS_CONNECTOR_API_KEY"`
		LaunchTimeoutSec int    `env:"PLUGIN_SERVERLESS_CONNECTOR_LAUNCH_TIMEOUT" default:"60"`
	}

	Signature struct {
		ForceVerify       bool     `env:"FORCE_VERIFYING_SIGNATURE" default:"false"`
		ThirdPartyEnabled bool     `env:"THIRD_PARTY_SIGNATURE_VERIFICATION_ENABLED" default:"false"`
		ThirdPartyKeys    []string `env:"THIRD_PARTY_SIGNATURE_VERIFICATION_PUBLIC_KEYS"`
	}
}

var AppCfg *AppConfig
