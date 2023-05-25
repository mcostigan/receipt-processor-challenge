package main

import (
	"encoding/json"
	_ "encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReceiptController struct {
	receiptService *ReceiptService
}

func NewReceiptController() *ReceiptController {
	return &ReceiptController{receiptService: NewReceiptService()}
}

func (controller *ReceiptController) handleProcessReceipts(c *gin.Context) {
	var newReceipt Receipt

	if err := json.NewDecoder(c.Request.Body).Decode(&newReceipt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	id := controller.receiptService.ProcessReceipt(&newReceipt)

	c.IndentedJSON(http.StatusOK, ProcessReceiptReturn{id})
}

func (controller *ReceiptController) handleGetPoints(c *gin.Context) {
	id := c.Param("receiptId")

	points, err := controller.receiptService.GetPoints(id)

	if err != nil {
		switch err.(type) {
		case *NoReceiptFoundError:
			c.IndentedJSON(http.StatusNotFound, err.Error())
			break
		default:
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong...")
		}

		return
	}

	c.IndentedJSON(http.StatusOK, GetPointsReturn{points})
}
