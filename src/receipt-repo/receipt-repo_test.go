package receipt_repo

import (
	"github.com/stretchr/testify/assert"
	"receipt-processor-challeng/src/model"
	"testing"
)

func TestInMemoryReceiptRepo_Get_ThrowsError(t *testing.T) {
	repo := NewInMemoryReceiptRepo()

	_, err := repo.Get("test")

	assert.Error(t, err)
}

func TestInMemoryReceiptRepo_Get_ReturnsReceipt(t *testing.T) {
	repo := InMemoryReceiptRepo{map[string]*model.Receipt{"test": {Id: "test"}}}

	r, err := repo.Get("test")

	assert.Equal(t, err, nil)
	assert.Equal(t, r.Id, "test")
}

func TestInMemoryReceiptRepo_Set_NewReceipt(t *testing.T) {
	repo := NewInMemoryReceiptRepo()

	r := repo.Set(&model.Receipt{Id: "test"})

	assert.Equal(t, r.Id, "test")
	assert.Equal(t, 1, len(repo.receipts))
}

func TestInMemoryReceiptRepo_Set_ExistingReceipt(t *testing.T) {
	repo := InMemoryReceiptRepo{map[string]*model.Receipt{"test": {Id: "test", Points: nil}}}

	r, _ := repo.Get("test")
	points := 5
	r.Points = &points

	repo.Set(r)

	r, _ = repo.Get("test")

	assert.Equal(t, r.Id, "test")
	assert.Equal(t, r.Points, &points)
	assert.Equal(t, 1, len(repo.receipts))
}
