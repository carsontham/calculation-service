package api

import (
	"log"
	"net/http"
)

func (s *APIServer) calculationHandler(w http.ResponseWriter, r *http.Request) error {
	//handle requests here
	log.Println("calculationHandler")
	return WriteJSON(w, http.StatusOK, nil)
}
