package api

import (
	"calculation-service/internal/calculation"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// this function calls to add all prices in an array and return the total price
func (s *APIServer) calculationHandler(w http.ResponseWriter, r *http.Request) error {
	//handle requests here
	log.Println("In calculationHandler")
	// logger, err := zap.NewProduction()

	// if err != nil {
	// 	log.Fatalf("can't initialize zap logger: %v", err)
	// }

	// logger.Info("hello")
	// logger.Error("error here")
	// Decode the request body into a struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	var req calculation.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding request body: ", err)
		//logger.Error("Error occured")

		return WriteJSON(w, http.StatusBadRequest, err)
	}

	// Calculate the total from the request.
	total, err := calculation.CalculateTotal(req)
	if err != nil {
		//logger.Info("error here")
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	// Return the total in a JSON response.
	return WriteJSON(w, http.StatusOK, map[string]float64{"result": total})
}

// this function calls to calculate the best discount based on current total price
// returns the best discount
func (s *APIServer) discountHandler(w http.ResponseWriter, r *http.Request) error {
	//handle requests here
	log.Println("In discountHandler")

	// Decode the request body into a struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	var req calculation.DiscountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		log.Println("Error decoding request body: ", err)
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	// Assert the types of the values in the vouchers
	for _, voucher := range req.Vouchers {
		_, ok := voucher["isPercentage"].(bool)
		if !ok {
			return WriteJSON(w, http.StatusBadRequest, fmt.Errorf("the request input is invalid"))
		}
		_, ok = voucher["value"].(float64)
		if !ok {
			return WriteJSON(w, http.StatusBadRequest, fmt.Errorf("the request input is invalid"))
		}
		_, ok = voucher["code"].(string)
		if !ok {
			return WriteJSON(w, http.StatusBadRequest, fmt.Errorf("the request input is invalid"))
		}
	}

	// Calculate the total from the request.
	res, err := calculation.GetBestVoucher(req)
	if err != nil {
		log.Println("Error calculating total: ", err)
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	// Return the total in a JSON response.
	return WriteJSON(w, http.StatusOK, res)
}
