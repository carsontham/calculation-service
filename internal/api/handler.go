package api

import (
	"calculation-service/internal/calculation"
	"encoding/json"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func (s *APIServer) calculationHandler(w http.ResponseWriter, r *http.Request) error {
	//handle requests here
	log.Println("In calculationHandler")
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	logger.Info("hello")
	// Decode the request body into a struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	var req calculation.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding request body: ", err)
		logger.Info("error")
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	// Calculate the total from the request.
	total, err := calculation.CalculateTotal(req)
	if err != nil {
		logger.Info("error here")
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	// Return the total in a JSON response.
	return WriteJSON(w, http.StatusOK, map[string]float64{"result": total})
}
