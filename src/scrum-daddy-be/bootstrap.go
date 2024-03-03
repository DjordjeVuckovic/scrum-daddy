package main

import (
	"net/http"
	"scrum-daddy-be/common/api"
	"scrum-daddy-be/pokerplanning"
)

func AddModules(server *api.Server) {
	server.AddRoute("/hello", helloHandler)
	addPokerPlanning(server)
}

func addPokerPlanning(server *api.Server) {
	pokerplanning.Main(server)
}

// @Summary Summarize your endpoint
// @Description Describe your endpoint
// @Tags v1
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} errors.ErrorResult
// @Router /hello [get]
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Your handler code here.
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		return
	}
}
