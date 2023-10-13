package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Version string
	Env     string
	Port    string
}

type APIError struct {
	Error string
}

func NewApiServer() *APIServer {
	return &APIServer{
		Version: "1.0.0",
		Env:     "Development",
		Port:    "8080",
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}
	return nil
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func (s *APIServer) Run() error {
	s.Port = os.Getenv("APP_PORT")
	s.Env = os.Getenv("ENV")
	if s.Env == "" {
		s.Env = "Dev"
	}
	log.Printf("Starting The Application In %s Mode On Port : %s.\n", s.Env, s.Port)
	r := mux.NewRouter()
	r.HandleFunc("/", s.HealthCheck)
	r.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))

	if err := http.ListenAndServe(":"+s.Port, r); err != nil {
		return err
	}
	return nil
}

func (s *APIServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("Error : Request Method Not Allowed ! "))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Application Is Running . . . "))
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		err := s.handleGetAccount(w, r)
		if err != nil {
			return err
		}
	}

	if r.Method == "POST" {
		err := s.handleCreateAccount(w, r)
		if err != nil {
			return err
		}
	}

	if r.Method == "DELETE" {
		err := s.handleDeleteAccount(w, r)
		if err != nil {
			return err
		}
	}
	// if r.Method == "PUT" {

	// }
	// if r.Method == "PATCH" {

	// }
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// func (s *APIServer) handleTransferMoney(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }
