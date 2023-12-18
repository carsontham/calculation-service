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
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	//	test := http.NewServeMux()

	// Endpoints				Method 		Function		Description
	// '/calculate'				POST 		calculate() 	- create new voucher
	router.HandleFunc("/calculate", makeHTTPHandleFunc(s.calculationHandler)).Methods(http.MethodGet, http.MethodPost)

	log.Println("JSON API server running on port:", s.listenAddr)
	//	http.ListenAndServe(s.listenAddr, router)
	log.Fatal(http.ListenAndServe(s.listenAddr, router))
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
