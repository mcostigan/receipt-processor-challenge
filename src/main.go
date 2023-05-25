package main

import (
	"receipt-processor-challeng/src/api"
)

func main() {
	router := api.GetRouter()
	router.Run("localhost:8080")
}
