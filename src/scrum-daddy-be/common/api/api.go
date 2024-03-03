package api

import (
	"encoding/json"
	"net/http"
	"scrum-daddy-be/common/errors"
)

type Server struct {
	listenAddress string
	mux           *http.ServeMux
}

func NewServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
		mux:           http.NewServeMux(),
	}
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

type apiFunc func(w http.ResponseWriter, r *http.Request) *errors.ErrorResult

func WriteJSON(w http.ResponseWriter, code int, body any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(body)
}

func MakeHandler(apiFunc apiFunc) http.HandlerFunc {
	return handleError(apiFunc)
}

func handleError(apiFunc apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(w, r); err != nil {
			switch err.Code {
			case http.StatusNotFound:
				_ = WriteJSON(w, http.StatusNotFound, err)
			case http.StatusBadRequest:
				_ = WriteJSON(w, http.StatusBadRequest, err)
			case http.StatusConflict:
				_ = WriteJSON(w, http.StatusConflict, err)
			default:
				_ = WriteJSON(w, http.StatusInternalServerError, err)
			}
		}
	}
}
