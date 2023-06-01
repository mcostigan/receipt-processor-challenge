package rules

import (
	"receipt-processor-challeng/src/model"
	"sync"
)

type ServiceInterface interface {
	PointReceipt(receipt *model.Receipt) int
}

type service struct {
	rules []Rule
}

func NewRulesService() ServiceInterface {
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
	pointsLock := sync.Mutex{}

	var wg sync.WaitGroup
	for _, rule := range service.rules {
		// evaluate each rule concurrently and update the total points
		wg.Add(1)
		go func(receipt *model.Receipt, rule Rule) {
			pointsLock.Lock()
			points += rule.evaluate(receipt)
			pointsLock.Unlock()
			wg.Done()
		}(receipt, rule)
	}

	// wait for the evaluation of all rules
	wg.Wait()
	return points
}
