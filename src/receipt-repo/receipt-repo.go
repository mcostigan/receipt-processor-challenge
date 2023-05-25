package receipt_repo

import (
	"fmt"
	"receipt-processor-challeng/src/model"
)

type ReceiptRepo interface {
	Get(id string) (*model.Receipt, error)
	Set(receipt *model.Receipt) *model.Receipt
}

type InMemoryReceiptRepo struct {
	receipts map[string]*model.Receipt
}

func NewInMemoryReceiptRepo() *InMemoryReceiptRepo {
	return &InMemoryReceiptRepo{map[string]*model.Receipt{}}
}

type NoReceiptFoundError struct {
	Id string
}

func (err *NoReceiptFoundError) Error() string {
	return fmt.Sprintf("Receipt with id %s does not exist", err.Id)
}

func (repo *InMemoryReceiptRepo) Get(id string) (*model.Receipt, error) {
	receipt, ok := repo.receipts[id]

	if !ok {
		return nil, &NoReceiptFoundError{id}
	}
	return receipt, nil
}

func (repo *InMemoryReceiptRepo) Set(receipt *model.Receipt) *model.Receipt {
	repo.receipts[receipt.Id] = receipt
	return receipt
}
