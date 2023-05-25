package main

import (
	"github.com/gin-gonic/gin"
	"receipt-processor-challeng/src/controller"
)

func main() {
	receiptController := controller.NewReceiptController()
	router := gin.Default()
	router.POST("/receipts/process", receiptController.HandleProcessReceipts)
	router.GET("/receipts/:receiptId/points", receiptController.HandleGetPoints)

	router.Run("localhost:8080")

}
