package main

import (
	// "log"
	"github.com/ArcherChu123/plugin-daemon/internal/server"
	"github.com/ArcherChu123/plugin-daemon/internal/types/app"
	"github.com/ArcherChu123/plugin-daemon/internal/utils/log"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	// "github.com/ArcherChu123/plugin-daemon/core/config"
)

func main() {

	// err := config.LoadFromEnv()
	// if err != nil {
	// 	log.Fatalf("加载配置失败: %v", err)
	// }
	var config app.Config

	// load env
	godotenv.Load()

	err := envconfig.Process("", &config)
	if err != nil {
		log.Panic("Error processing environment variables: %s", err.Error())
	}

	config.SetDefault()

	if err := config.Validate(); err != nil {
		log.Panic("Invalid configuration: %s", err.Error())
	}

	(&server.App{}).Run(&config)
}
