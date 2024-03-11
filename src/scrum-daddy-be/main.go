package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/common/swagger"
	"scrum-daddy-be/identity"
	"scrum-daddy-be/pokerplanning"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("No .env file found")
		panic(err)
	}
	port := os.Getenv("PORT")
	server := api.NewServer(":" + port)

	dbConnection := db.Connect()
	defer dbConnection.Close()

	AddModules(server, dbConnection)

	swagger.SetupSwagger(server.GetMux())
	slog.Info("Starting server on port %s\n", port)
	if err := server.Start(); err != nil {
		slog.Error("Could not start server: %s\n", err)
	}
}

func AddModules(s *api.Server, db *db.Database) {
	moduleContainers := CreateModuleContainers(s, db)

	pokerplanning.Main(moduleContainers.PokerPlanning)

	identity.Main(moduleContainers.Identity)
}
