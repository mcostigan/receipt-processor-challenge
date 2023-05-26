package rules

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"os"
	"receipt-processor-challeng/src/model"
	"testing"
)

func TestReceiptPointingExample1(t *testing.T) {
	receiptJson, _ := os.Open("../../examples/example-receipt-1.json")
	defer func(receiptJson *os.File) {
		err := receiptJson.Close()
		if err != nil {

		}
	}(receiptJson)

	var exampleReceipt model.Receipt
	err := json.NewDecoder(receiptJson).Decode(&exampleReceipt)
	if err != nil {
		return
	}

	service := NewRulesService()
	points := service.PointReceipt(&exampleReceipt)

	assert.Equal(t, 28, points)
}

func TestReceiptPointingExample2(t *testing.T) {
	receiptJson, _ := os.Open("../../examples/example-receipt-2.json")
	defer func(receiptJson *os.File) {
		err := receiptJson.Close()
		if err != nil {

		}
	}(receiptJson)

	var exampleReceipt model.Receipt
	err := json.NewDecoder(receiptJson).Decode(&exampleReceipt)
	if err != nil {
		return
	}

	service := NewRulesService()
	points := service.PointReceipt(&exampleReceipt)

	assert.Equal(t, 109, points)
}
