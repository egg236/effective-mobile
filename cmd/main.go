package main

import (
	"effective-mobile/api"
	"effective-mobile/app"
	"effective-mobile/config"
	"log"
	"log/slog"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load configuration: ", err)
	}

	app, err := app.NewApp()
	if err != nil {
		log.Fatal("failed to initialize application: ", err)
	}

	slog.Info("application started")

	server := api.NewServer(app)

	slog.Info("starting server on port address", "address", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("fatal error occured while executing the server", "error", err.Error())
		log.Fatal("failed to start server: ", err)
	}

}
