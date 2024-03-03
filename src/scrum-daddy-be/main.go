package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/common/swagger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		panic(err)
	}
	port := os.Getenv("PORT")
	server := api.NewServer(":" + port)
	AddModules(server)

	swagger.SetupSwagger(server.GetMux())
	log.Printf("Starting server on port %s\n", port)
	if err := server.Start(); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
