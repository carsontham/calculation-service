package calculation

import (
	"calculation-service/internal/middleware"
	"errors"
	"fmt"
	"log"
	"time"
)

type CalculationRequest struct {
	Items        []map[string]interface{} `json:"items"`
	Operation    string                   `json:"operation"`
	DiscountRate float64                  `json:"discount_rate"`
}

type CalculationResponse struct {
	Total float64 `json:"total"`
}

type DiscountRequest struct {
	Vouchers   []map[string]interface{} `json:"vouchers"`
	TotalPrice float64                  `json:"total_price"`
}

type DiscountResponse struct {
	VoucherCode   string  `json:"voucher_code"`
	NewTotalPrice float64 `json:"net_price"`
	AmountSaved   float64 `json:"amt_saved"`
}

func CalculateTotal(req CalculationRequest) (float64, error) {
	logger := middleware.NewZapLogger()
	logger.Info("In CalculateTotal")
	total := 0.0
	for _, item := range req.Items {
		for key, value := range item {
			itemData, ok := value.(map[string]interface{})
			if !ok {
				return 0, fmt.Errorf("item does not contain valid data")
			}

			price, ok := itemData["price"].(float64)
			if !ok {
				return 0, fmt.Errorf("item data does not contain a valid price")
			}

			qty, ok := itemData["qty"].(float64)
			if !ok {
				return 0, fmt.Errorf("item data does not contain a valid qty")
			}

			// Calculate total for the current item
			itemTotal := price * qty

			// Add to the overall total
			total += itemTotal

			// Print the variable name (e.g., "Car")
			fmt.Println("Variable Name:", key)
		}
	}

	return total, nil
}

func GetBestVoucher(req DiscountRequest) (DiscountResponse, error) {
	start := time.Now()
	vouchers := req.Vouchers
	total := req.TotalPrice
	if len(req.Vouchers) == 0 {
		return DiscountResponse{}, errors.New("no vouchers provided")
	}

	bestVoucher := vouchers[0]
	minNetPrice, err := calculateNetPrice(total, bestVoucher)

	if err != nil {
		return DiscountResponse{}, err
	}

	for i, voucher := range vouchers[1:] {

		log.Println("entered loop", i)
		netPrice, err := calculateNetPrice(total, voucher)

		if err != nil {
			return DiscountResponse{}, err
		}

		if netPrice < minNetPrice {
			bestVoucher = voucher
			minNetPrice = netPrice
		}
	}

	amtSaved := req.TotalPrice - minNetPrice
	log.Println("time taken", time.Since(start))
	return DiscountResponse{
		VoucherCode:   bestVoucher["code"].(string),
		NewTotalPrice: minNetPrice,
		AmountSaved:   amtSaved}, nil
}

func calculateNetPrice(total float64, voucher map[string]interface{}) (float64, error) {
	log.Println("entered calculateNetPrice()")
	if voucher["isPercentage"].(bool) {
		return total - total*(voucher["value"].(float64)/100), nil
	}
	log.Println("here", total, voucher["value"].(float64), voucher["code"].(string))
	return total - voucher["value"].(float64), nil
}
