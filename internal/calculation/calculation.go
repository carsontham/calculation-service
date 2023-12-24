package calculation

import (
	"log"
	"sync"
)

type CalculationRequest struct {
	Items        []map[string]interface{} `json:"items"`
	Operation    string                   `json:"operation"`
	DiscountRate float64                  `json:"discount_rate"`
}

type CalculationResponse struct {
	Total float64 `json:"total"`
}

func CalculateTotal(req CalculationRequest) (float64, error) {
	results := make(chan float64, len(req.Items))
	var wg sync.WaitGroup

	for _, item := range req.Items {
		wg.Add(1)
		go func(item map[string]interface{}) {
			defer wg.Done()
			price, ok := item["price"].(float64)
			if !ok {
				log.Println("item does not have a valid price")
				return
			}
			results <- price
		}(item)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0.0
	for price := range results {
		total += price
	}
	return total, nil
}
