package api

import (
	"encoding/json"
	"net/http"
	"os"
	"scrum-daddy-be/common/results"
	"strings"
)

type Server struct {
	listenAddress string
	mux           *http.ServeMux
}

type ServerOption func(*Server)

func NewServer(listenAddress string, opts ...ServerOption) *Server {
	server := &Server{
		listenAddress: listenAddress,
		mux:           http.NewServeMux(),
	}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.listenAddress, s.mux)
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) GetMux() *http.ServeMux {
	return s.mux
}

func (s *Server) AddRoute(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler)
}
func WithCORS() ServerOption {
	return func(s *Server) {
		s.mux.Handle("/", enableCORS(s.mux))
	}
}

func enableCORS(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ORIGIN"))
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Api-Key")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	}
}

type apiFunc func(w http.ResponseWriter, r *http.Request) *results.ErrorResult

func WriteJSON(w http.ResponseWriter, code int, body any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(body)
}

func writeErrorJSON(w http.ResponseWriter, code int, body any) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(body)
}

func MakeHandler(apiFunc apiFunc) http.HandlerFunc {
	return enableCORS(handleError(apiFunc))
}

func handleError(apiFunc apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(w, r); err != nil {
			switch err.Code {
			case http.StatusNotFound:
				_ = writeErrorJSON(w, http.StatusNotFound, err)
			case http.StatusBadRequest:
				_ = writeErrorJSON(w, http.StatusBadRequest, err)
			case http.StatusConflict:
				_ = writeErrorJSON(w, http.StatusConflict, err)
			default:
				_ = writeErrorJSON(w, http.StatusInternalServerError, err)
			}
		}
	}
}

func WithApiKeyAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	apiKey := os.Getenv("API_KEY_AUTH")
	return func(w http.ResponseWriter, r *http.Request) {
		apiKeyHeader := r.Header.Get("X-Api-Key")
		if apiKeyHeader == "" {
			noKeyErr := results.NewErrorResult(
				http.StatusUnauthorized,
				"Api Key Not found",
				"Api key is missing in the request header.")
			_ = WriteJSON(w, http.StatusUnauthorized, noKeyErr)
			return
		}

		if strings.Compare(apiKeyHeader, apiKey) != 0 {
			badKeyErr := results.NewErrorResult(http.StatusForbidden, "Api Key not valid", "")
			_ = WriteJSON(w, http.StatusUnauthorized, badKeyErr)
			return
		}

		handlerFunc(w, r)
	}

}
