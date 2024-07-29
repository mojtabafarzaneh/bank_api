package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type ApiServer struct {
	listenAdder string
}

func NewAPIServer(listenAdder string) *ApiServer {
	return &ApiServer{
		listenAdder: listenAdder,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHttpHandleFunc(s.HandleAccount))

	fmt.Println("server is running on the port: ", s.listenAdder)

	http.ListenAndServe(s.listenAdder, router)

}

func (s *ApiServer) HandleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.HandleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.HandleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.HandleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *ApiServer) HandleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("mojtaba", "farzaneh")

	return WriteJSON(w, http.StatusOK, account)
}

func (s *ApiServer) HandleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) HandleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s ApiServer) HandleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
