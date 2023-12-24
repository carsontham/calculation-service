package api

import (
	"calculation-service/internal/middleware"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	router     *mux.Router
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		router:     mux.NewRouter(),
	}
}

func (s *APIServer) Run() {
	s.router.Use(middleware.LoggingMiddleware)
	s.SetupRoutes()

	log.Println("JSON API server running on port:", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, s.router))
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle error
			WriteJSON(w, http.StatusBadRequest, err)
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
