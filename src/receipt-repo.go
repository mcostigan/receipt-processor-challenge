package main

import "fmt"

type ReceiptRepo interface {
	Get(id string) (*Receipt, error)
	Set(receipt *Receipt) *Receipt
}

type InMemoryReceiptRepo struct {
	receipts map[string]*Receipt
}

func NewInMemoryReceiptRepo() *InMemoryReceiptRepo {
	return &InMemoryReceiptRepo{map[string]*Receipt{}}
}

type NoReceiptFoundError struct {
	id string
}

func (err *NoReceiptFoundError) Error() string {
	return fmt.Sprintf("Receipt with id %s does not exist", err.id)
}

func (repo *InMemoryReceiptRepo) Get(id string) (*Receipt, error) {
	receipt, ok := repo.receipts[id]

	if !ok {
		return nil, &NoReceiptFoundError{id}
	}
	return receipt, nil
}

func (repo *InMemoryReceiptRepo) Set(receipt *Receipt) *Receipt {
	repo.receipts[receipt.Id] = receipt
	return receipt
}
