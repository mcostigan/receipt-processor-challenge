package api

import (
	"github.com/gin-gonic/gin"
	receiptController "receipt-processor-challeng/src/receipt-controller"
)

// GetRouter Creates a default gin router and maps endpoints to handler funcs
func GetRouter() *gin.Engine {
	controller := receiptController.NewReceiptController()
	router := gin.Default()
	router.POST("/receipts/process", controller.HandleProcessReceipts)
	router.GET("/receipts/:receiptId/points", controller.HandleGetPoints)
	return router
}
