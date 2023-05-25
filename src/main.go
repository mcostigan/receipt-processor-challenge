package main

import "github.com/gin-gonic/gin"

func main() {
	receiptController := NewReceiptController()
	router := gin.Default()
	router.POST("/receipts/process", receiptController.handleProcessReceipts)
	router.GET("/receipts/:receiptId/points", receiptController.handleGetPoints)

	router.Run("localhost:8080")

}
