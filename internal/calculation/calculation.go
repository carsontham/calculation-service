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
	//start := time.Now()
	total := 0.0
	for _, i := range req.Items {
		price, ok := i["price"].(float64)
		if !ok {
			return 0, fmt.Errorf("price is not a float64")
		}
		total += price
		//log.Println("index: ", index)
	}

	//log.Println("Total time using for loop:  ", time.Since(start))
	return total, nil
}
