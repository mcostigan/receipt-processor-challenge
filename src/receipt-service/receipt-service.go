package receipt_service

import (
	"github.com/google/uuid"
	"receipt-processor-challeng/src/model"
	"receipt-processor-challeng/src/receipt-repo"
	"receipt-processor-challeng/src/rules"
)

type ReceiptService interface {
	ProcessReceipt(receipt *model.Receipt) string
	GetPoints(id string) (int, error)
}
type service struct {
	rulesService rules.ServiceInterface
	receiptRepo  receipt_repo.ReceiptRepo
}

func NewReceiptService() ReceiptService {
	return &service{
		rulesService: rules.NewRulesService(),
		receiptRepo:  receipt_repo.NewInMemoryReceiptRepo(),
	}
}

// ProcessReceipt Assigns a unique ID to the receipt and persists (via the repository)
func (service *service) ProcessReceipt(receipt *model.Receipt) string {
	id := uuid.NewString()
	receipt.Id = id
	service.receiptRepo.Set(receipt)
	return id
}

// GetPoints Returns the number of points earned on a receipt.
// If the receipt has never been queried before, the amount is calculated and stored with the receipt.
// Else, the value is fetched via.
func (service *service) GetPoints(id string) (int, error) {
	receipt, err := service.receiptRepo.Get(id)
	if err != nil {
		return -1, err
	}
	// Only calculate the points one time
	// After that, persist with the receipt object
	if receipt.Points == nil {
		points := service.rulesService.PointReceipt(receipt)
		receipt.Points = &points

		// This part isn't necessary with an in-memory data store
		// But if the repo connected to another database, we would want to call a `save` procedure
		service.receiptRepo.Set(receipt)
	}

	return *receipt.Points, nil
}
