package api

import (
	"github.com/gin-gonic/gin"
	receiptController "receipt-processor-challeng/src/receipt-controller"
)

func RunApi() {
	controller := receiptController.NewReceiptController()
	router := gin.Default()
	router.POST("/receipts/process", controller.HandleProcessReceipts)
	router.GET("/receipts/:receiptId/points", controller.HandleGetPoints)

	router.Run("localhost:8080")
}
