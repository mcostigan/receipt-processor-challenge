package main

type RulesServiceInterface interface {
	PointReceipt(receipt *Receipt) int
}

type RulesService struct {
	rules []Rule
}

func NewRulesService() *RulesService {
	return &RulesService{[]Rule{
		&RetailerLengthRule{},
		&RoundDollarAmountRule{},
		&PointTwoFiveRule{},
		&ItemsLengthRule{},
		&TrimmedItemDescriptionRule{},
		&OddDayRule{},
		&TwoPMTo4PMRule{}}}
}

func (service *RulesService) PointReceipt(receipt *Receipt) int {
	points := 0
	for _, rule := range service.rules {
		points += rule.evaluate(receipt)
	}
	return points
}
