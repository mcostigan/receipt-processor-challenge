package main

import (
	"receipt-processor-challeng/src/api"
)

func main() {
	router := api.GetRouter()
	if err := router.Run("0.0.0.0:8080"); err != nil {
		return
	}
}
