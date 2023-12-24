package main

import "calculation-service/internal/api"

func main() {
	calculationServer := api.NewAPIServer(":4000")
	calculationServer.Run()
}
