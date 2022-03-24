package main

import (
	"apiGetaway/config"
	"apiGetaway/package/logger"
	"apiGetaway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "admin_api_gateway")

	grpcClients, _ := services.NewGrpcClients(&cfg)

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      cfg,
		Services: grpcClients,
	})

	server.Run(cfg.HttpPort)
}
