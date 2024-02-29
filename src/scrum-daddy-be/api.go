package main

import "net/http"

func BindEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/hello", helloHandler)
}

// @Summary Summarize your endpoint
// @Description Describe your endpoint
// @Tags v1
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} errors.ApiError
// @Router /hello [get]
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Your handler code here.
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		return
	}
}
