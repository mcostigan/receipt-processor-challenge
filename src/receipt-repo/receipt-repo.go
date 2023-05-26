package receipt_repo

import (
	"fmt"
	"receipt-processor-challeng/src/model"
	"sync"
)

type ReceiptRepo interface {
	Get(id string) (*model.Receipt, error)
	Set(receipt *model.Receipt) *model.Receipt
}

type inMemoryReceiptRepo struct {
	receipts map[string]*model.Receipt
}

func NewInMemoryReceiptRepo() *inMemoryReceiptRepo {
	return &inMemoryReceiptRepo{map[string]*model.Receipt{}}
}

type NoReceiptFoundError struct {
	Id string
}

func (err *NoReceiptFoundError) Error() string {
	return fmt.Sprintf("Receipt with id %s does not exist", err.Id)
}

var mutex = &sync.RWMutex{}

// Get Searches the data store for a receipt with a given id and returns.
// Throws an exception if no such id exists
func (repo *inMemoryReceiptRepo) Get(id string) (*model.Receipt, error) {
	mutex.Lock()
	receipt, ok := repo.receipts[id]
	mutex.Unlock()

	if !ok {
		return nil, &NoReceiptFoundError{id}
	}
	return receipt, nil
}

// Set Saves an instance of receipt, keyed on its id
func (repo *inMemoryReceiptRepo) Set(receipt *model.Receipt) *model.Receipt {
	mutex.Lock()
	repo.receipts[receipt.Id] = receipt
	mutex.Unlock()
	return receipt
}
