package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	_ "scrum-daddy-be/common/errors"
	"scrum-daddy-be/common/swagger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		panic(err)
	}

	mux := http.NewServeMux()
	BindEndpoints(mux)
	swagger.SetupSwagger(mux)

	port := os.Getenv("PORT")
	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
