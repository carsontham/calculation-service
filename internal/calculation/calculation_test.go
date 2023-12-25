package calculation

import "testing"

// Table driven tests allows for clarity and readability
func TestCalculateTotal(t *testing.T) {
	tests := map[string]struct {
		input   CalculationRequest
		want    float64
		wantErr bool
	}{
		"different objects": {input: CalculationRequest{Items: []map[string]interface{}{{"Car": map[string]interface{}{"Make": "Toyota", "Model": "Camry", "Year": 2022, "price": 10.0, "qty": 1.0}}, {"Products": map[string]interface{}{"Name": "Vegetables", "ItemCode": "TEST101", "price": 20.0, "qty": 3.0}}}}, want: 70.0},
		"simple":            {input: CalculationRequest{Items: []map[string]interface{}{{"Car": map[string]interface{}{"Make": "Toyota", "Model": "Camry", "Year": 2022, "price": 10.0, "qty": 1.0}}, {"Car": map[string]interface{}{"Make": "Honda", "Model": "Civic", "Year": 2023, "price": 20.0, "qty": 1.0}}}}, want: 30.0},
		"int":               {input: CalculationRequest{Items: []map[string]interface{}{{"Car": map[string]interface{}{"Make": "Toyota", "Model": "Camry", "Year": 2022, "price": 10, "qty": 1.0}}, {"Car": map[string]interface{}{"Make": "Honda", "Model": "Civic", "Year": 2023, "price": 20, "qty": 1}}}}, wantErr: true},
		"invalid price":     {input: CalculationRequest{Items: []map[string]interface{}{{"Car": map[string]interface{}{"Make": "Toyota", "Model": "Camry", "Year": 2022, "price": "String", "qty": 1.0}}, {"Car": map[string]interface{}{"Make": "Honda", "Model": "Civic", "Year": 2023, "price": "string", "qty": 1.0}}}}, wantErr: true},
		"empty":             {input: CalculationRequest{Items: []map[string]interface{}{{"Car": map[string]interface{}{"Make": "Toyota", "Model": "Camry", "Year": 2022}}, {"Car": map[string]interface{}{"Make": "Honda", "Model": "Civic", "Year": 2023}}}}, wantErr: true},
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
