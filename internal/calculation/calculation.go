package calculation

import "fmt"

type CalculationRequest struct {
	Items        []map[string]interface{} `json:"items"`
	Operation    string                   `json:"operation"`
	DiscountRate float64                  `json:"discount_rate"`
}

type CalculationResponse struct {
	Total float64 `json:"total"`
}

func CalculateTotal(req CalculationRequest) (float64, error) {
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
