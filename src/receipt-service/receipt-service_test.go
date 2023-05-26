package receipt_service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"receipt-processor-challeng/src/model"
	"receipt-processor-challeng/src/receipt-repo"
	"testing"
)

type MockRepo struct {
	mock.Mock
}

func (r *MockRepo) Get(id string) (*model.Receipt, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Receipt), args.Error(1)
}

func (r *MockRepo) Set(receipt *model.Receipt) *model.Receipt {
	args := r.Called(receipt)
	return args.Get(0).(*model.Receipt)
}

type MockRulesService struct {
	mock.Mock
}

func (service *MockRulesService) PointReceipt(r *model.Receipt) int {
	args := service.Called(r)
	return args.Get(0).(int)
}

func TestReceiptService_ProcessReceipt(t *testing.T) {
	mockRepo := &MockRepo{}
	mockRulesService := &MockRulesService{}

	service := &service{mockRulesService, mockRepo}

	receipt := &model.Receipt{}
	mockRepo.On("Set", receipt).Return(receipt)

	service.ProcessReceipt(receipt)

	mockRepo.AssertCalled(t, "Set", receipt)
	assert.NotEqual(t, receipt.Id, "")

}

func TestReceiptService_GetPoints_FirstTime(t *testing.T) {
	mockRepo := &MockRepo{}
	mockRulesService := &MockRulesService{}

	service := &service{mockRulesService, mockRepo}

	// return receipt with nil points
	receipt := &model.Receipt{}
	mockRepo.On("Get", "test").Return(receipt, nil)
	mockRepo.On("Set", receipt).Return(receipt)

	mockRulesService.On("PointReceipt", receipt).Return(123)

	points, _ := service.GetPoints("test")

	assert.Equal(t, points, 123)
	mockRepo.AssertCalled(t, "Get", "test")
	mockRepo.AssertCalled(t, "Set", receipt)
	mockRulesService.AssertCalled(t, "PointReceipt", receipt)
}

func TestReceiptService_GetPoints_Subsequent(t *testing.T) {
	mockRepo := &MockRepo{}
	mockRulesService := &MockRulesService{}

	service := &service{mockRulesService, mockRepo}

	initialPoints := 50
	// return receipt with nil points
	receipt := &model.Receipt{Points: &initialPoints}
	mockRepo.On("Get", "test").Return(receipt, nil)
	mockRepo.On("Set", receipt).Return(receipt)

	mockRulesService.On("PointReceipt", receipt).Return(123)

	points, _ := service.GetPoints("test")

	assert.Equal(t, points, 50)
	mockRepo.AssertCalled(t, "Get", "test")
	mockRepo.AssertNotCalled(t, "Set", receipt)
	mockRulesService.AssertNotCalled(t, "PointReceipt", receipt)
}

func TestReceiptService_GetPoints_BadId(t *testing.T) {
	mockRepo := &MockRepo{}
	mockRulesService := &MockRulesService{}

	service := &service{mockRulesService, mockRepo}

	mockRepo.On("Get", "test").Return(nil, &receipt_repo.NoReceiptFoundError{Id: "test"})

	_, err := service.GetPoints("test")

	assert.Error(t, err)
	mockRepo.AssertCalled(t, "Get", "test")

}
