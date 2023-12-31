package api

import (
	"calculation-service/internal/calculation"
	"calculation-service/internal/middleware"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// this function calls to add all prices in an array and return the total price
func (s *APIServer) calculationHandler(w http.ResponseWriter, r *http.Request) error {
	logger := middleware.NewZapLogger()
	logger.Info("in calculationHandler")

	// Decode the request body into a struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	var req calculation.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("error decoding request body", zap.Error(err))

		return WriteJSON(w, http.StatusBadRequest, err)
	}
	// Calculate the total from the request.
	total, err := calculation.CalculateTotal(req)
	if err != nil {
		logger.Error("error calculating total", zap.Error(err))
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	// Return the total in a JSON response.
	logger.Info("calculate total complete")
	return WriteJSON(w, http.StatusOK, map[string]float64{"result": total})
}

// this function calls to calculate the best discount based on current total price
// returns the best discount
func (s *APIServer) discountHandler(w http.ResponseWriter, r *http.Request) error {
	logger := middleware.NewZapLogger()
	logger.Info("in discountHandler")

	// Decode the request body into a struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	var req calculation.DiscountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("error decoding request body", zap.Error(err))
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	// Assert the types of the values in the vouchers
	for _, voucher := range req.Vouchers {
		_, ok := voucher["isPercentage"].(bool)
		if !ok {
			logger.Warn("the request input is invalid")
			return WriteJSON(w, http.StatusBadRequest, fmt.Errorf("the request input is invalid"))
		}
		_, ok = voucher["value"].(float64)
		if !ok {
			logger.Warn("the request input is invalid")
			return WriteJSON(w, http.StatusBadRequest, fmt.Errorf("the request input is invalid"))
		}
		_, ok = voucher["code"].(string)
		if !ok {
			logger.Warn("the request input is invalid")
			return WriteJSON(w, http.StatusBadRequest, fmt.Errorf("the request input is invalid"))
		}
	}
	logger.Info("request input is valid")
	// Calculate the total from the request.
	res, err := calculation.GetBestVoucher(req)
	if err != nil {
		logger.Error("error calculating total", zap.Error(err))
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	logger.Info("calculate total complete")
	// Return the total in a JSON response.
	return WriteJSON(w, http.StatusOK, res)
}
