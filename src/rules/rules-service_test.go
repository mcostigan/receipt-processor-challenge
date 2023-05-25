package rules

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"os"
	"receipt-processor-challeng/src/model"
	"testing"
)

func TestReceiptPointingExample1(t *testing.T) {
	receiptJson, _ := os.Open("../examples/example-receipt-1.json")
	defer receiptJson.Close()

	var exampleReceipt model.Receipt
	json.NewDecoder(receiptJson).Decode(&exampleReceipt)

	service := NewRulesService()
	points := service.PointReceipt(&exampleReceipt)

	assert.Equal(t, 28, points)
}

func TestReceiptPointingExample2(t *testing.T) {
	receiptJson, _ := os.Open("../examples/example-receipt-2.json")
	defer receiptJson.Close()

	var exampleReceipt model.Receipt
	json.NewDecoder(receiptJson).Decode(&exampleReceipt)

	service := NewRulesService()
	points := service.PointReceipt(&exampleReceipt)

	assert.Equal(t, 109, points)
}
