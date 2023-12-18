package main

import "calculation-service/api"

func main() {
	calculationServer := api.NewAPIServer(":5000")
	calculationServer.Run()
}
