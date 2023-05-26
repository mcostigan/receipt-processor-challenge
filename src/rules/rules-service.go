package rules

import (
	"receipt-processor-challeng/src/model"
)

type ServiceInterface interface {
	PointReceipt(receipt *model.Receipt) int
}

type service struct {
	rules []Rule
}

func NewRulesService() *service {
	return &service{[]Rule{
		&RetailerLengthRule{},
		&RoundDollarAmountRule{},
		&PointTwoFiveRule{},
		&ItemsLengthRule{},
		&TrimmedItemDescriptionRule{},
		&OddDayRule{},
		&TwoPMTo4PMRule{}}}
}

// PointReceipt Find the total points earned on the receipt /**
func (service *service) PointReceipt(receipt *model.Receipt) int {
	points := 0
	for _, rule := range service.rules {
		points += rule.evaluate(receipt)
	}
	return points
}
