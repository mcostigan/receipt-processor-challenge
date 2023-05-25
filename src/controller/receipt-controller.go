package controller

import (
	"encoding/json"
	_ "encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"receipt-processor-challeng/src/model"
	"receipt-processor-challeng/src/receipt-repo"
	"receipt-processor-challeng/src/receipt-service"
)

type ReceiptController struct {
	receiptService *receipt_service.ReceiptService
}

func NewReceiptController() *ReceiptController {
	return &ReceiptController{receiptService: receipt_service.NewReceiptService()}
}

func (controller *ReceiptController) HandleProcessReceipts(c *gin.Context) {
	var newReceipt model.Receipt

	if err := json.NewDecoder(c.Request.Body).Decode(&newReceipt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	id := controller.receiptService.ProcessReceipt(&newReceipt)

	c.IndentedJSON(http.StatusOK, model.ProcessReceiptReturn{id})
}

func (controller *ReceiptController) HandleGetPoints(c *gin.Context) {
	id := c.Param("receiptId")

	points, err := controller.receiptService.GetPoints(id)

	if err != nil {
		switch err.(type) {
		case *receipt_repo.NoReceiptFoundError:
			c.IndentedJSON(http.StatusNotFound, err.Error())
			break
		default:
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong...")
		}

		return
	}

	c.IndentedJSON(http.StatusOK, model.GetPointsReturn{points})
}
