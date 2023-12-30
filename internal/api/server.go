package api

import (
	"calculation-service/internal/middleware"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
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
	logger := middleware.NewZapLogger()

	s.router.Use(middleware.LoggingMiddleware)
	s.SetupRoutes()

	logger.Info("calculation-server running... ", zap.String("port", s.listenAddr))
	//log.Println("JSON API server running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, s.router)
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
