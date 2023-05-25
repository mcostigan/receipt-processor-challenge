package main

import "github.com/google/uuid"

type ReceiptService struct {
	rulesService RulesServiceInterface
	receiptRepo  ReceiptRepo
}

func NewReceiptService() *ReceiptService {
	return &ReceiptService{
		rulesService: NewRulesService(),
		receiptRepo:  NewInMemoryReceiptRepo(),
	}
}

// ProcessReceipt Assigns a unique ID to the receipt and persists (via the repository)
func (service *ReceiptService) ProcessReceipt(receipt *Receipt) string {
	id := uuid.NewString()
	receipt.Id = id
	service.receiptRepo.Set(receipt)
	return id
}

func (service *ReceiptService) GetPoints(id string) (int, error) {
	receipt, err := service.receiptRepo.Get(id)
	if err != nil {
		return -1, err
	}
	// Only calculate the points one time
	// After that, persist with the receipt object
	if receipt.points == nil {
		points := service.rulesService.PointReceipt(receipt)
		receipt.points = &points

		// This part isn't necessary with an in-memory data store
		// But if the repo connected to another database, we would want to call a `save` procedure
		service.receiptRepo.Set(receipt)
	}

	return *receipt.points, nil
}
