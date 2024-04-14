package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/common/logger"
	"scrum-daddy-be/common/swagger"
	"scrum-daddy-be/identity"
	"scrum-daddy-be/pokerplanning"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Api-Key

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("No .env file found")
		panic(err)
	}
	port := os.Getenv("PORT")
	server := api.NewServer(":" + port)

	dbConnection := db.Connect()
	defer dbConnection.Close()

	logger.ConfigureLogger()

	AddModules(server, dbConnection)

	swagger.SetupSwagger(server.GetMux())

	slog.Info("Starting server on port", "PORT", port)
	if err := server.Start(); err != nil {
		slog.Error("Could not start server", "error", err)
	}
}

func AddModules(s *api.Server, db *db.Database) {
	moduleContainers := CreateModules(s, db)

	pokerplanning.Main(moduleContainers.PokerPlanning)

	identity.Main(moduleContainers.Identity)
}
