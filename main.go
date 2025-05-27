package main

import (
	"Mars/database"
	"Mars/database/controller"
	"Mars/server"
	"Mars/shared/configure"
	"Mars/shared/utils/json"
	"Mars/shared/utils/json/json_se/sonic"
	"log/slog"
)

func main() {
	slog.Info("Starting application....")
	configure.NewConfig()
	json.New()
	sonic.PreTouch()
	database.Database()
	controller.RemoveAllBundlerBuilds()
	server.Server()
}
