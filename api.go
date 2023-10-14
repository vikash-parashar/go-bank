package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Version string
	Env     string
	Port    string
	DB      *sql.DB
	Store   Storage
}

func NewApiServer(store Storage) *APIServer {
	return &APIServer{
		Version: "1.0.0",
		Env:     "Development",
		Port:    "8080",
		Store:   store,
	}
}

type APIError struct {
	Error string
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
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)

	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("request method not allowed : %s ", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	Id := mux.Vars(r)["id"]
	// account := NewAccount("vikash", "parashar")
	nid, _ := strconv.Atoi(Id)
	return writeJSON(w, http.StatusOK, Account{ID: nid})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
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
	r.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleGetAccount))
	if err := http.ListenAndServe(":"+s.Port, r); err != nil {
		return err
	}
	return nil
}
