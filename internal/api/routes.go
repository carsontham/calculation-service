package api

import "net/http"

func (s *APIServer) SetupRoutes() {
	// Endpoints				Method 		Function		Description
	// '/calculate'				POST 		calculate() 	- create new voucher
	s.router.HandleFunc("/calculate", makeHTTPHandleFunc(s.calculationHandler)).Methods(http.MethodGet, http.MethodPost)
	s.router.HandleFunc("/discount", makeHTTPHandleFunc(s.discountHandler)).Methods(http.MethodGet, http.MethodPost)
}
