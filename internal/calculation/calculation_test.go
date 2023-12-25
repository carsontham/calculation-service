package calculation

import "testing"

// Table driven tests allows for clarity and readability
func TestCalculateTotal(t *testing.T) {
	tests := map[string]struct {
		input   CalculationRequest
		want    float64
		wantErr bool
	}{
		"simple":        {input: CalculationRequest{Items: []map[string]interface{}{{"price": 10.0}, {"price": 20.0}, {"price": 30.0}}}, want: 60.0},
		"int":           {input: CalculationRequest{Items: []map[string]interface{}{{"price": 10}, {"price": 20}, {"price": 30}}}, wantErr: true},
		"invalid price": {input: CalculationRequest{Items: []map[string]interface{}{{"price": "invalid"}}}, wantErr: true},
		"empty":         {input: CalculationRequest{Items: []map[string]interface{}{}}, want: 0.0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := CalculateTotal(tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Test: %v, CalculateTotal error = %v, wantErr %v", name, err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("Test: %v, CalculateTotal() = %v, want %v", name, got, tc.want)
			}
		})
	}
}
