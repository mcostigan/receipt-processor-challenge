package rules

import (
	"math"
	"receipt-processor-challeng/src/model"
	"regexp"
	_ "regexp"
	"strings"
)

type Rule interface {
	// evaluates a rule against a given receipt and returns the score
	evaluate(*model.Receipt) int
}

// RetailerLengthRule One point for every alphanumeric character in the retailer name.
type RetailerLengthRule struct{}

func (rule *RetailerLengthRule) evaluate(receipt *model.Receipt) int {
	return len(regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(receipt.Retailer, ""))
}

// RoundDollarAmountRule 50 points if the total is a round dollar amount with no cents.
type RoundDollarAmountRule struct{}

func (rule *RoundDollarAmountRule) evaluate(receipt *model.Receipt) int {
	points := 0
	if receipt.Total.Cents%100 == 0 {
		points = 50
	}
	return points
}

// PointTwoFiveRule 25 points if the total is a multiple of 0.25.
type PointTwoFiveRule struct{}

func (rule *PointTwoFiveRule) evaluate(receipt *model.Receipt) int {
	points := 0
	if receipt.Total.Cents%25 == 0 {
		points = 25
	}
	return points
}

// ItemsLengthRule 5 points for every two items on the receipt.
type ItemsLengthRule struct{}

func (rule *ItemsLengthRule) evaluate(receipt *model.Receipt) int {
	// perform integer division to get pairs and multiply by 5
	return len(receipt.Items) / 2 * 5
}

// TrimmedItemDescriptionRule If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
type TrimmedItemDescriptionRule struct{}

func (rule *TrimmedItemDescriptionRule) evaluate(receipt *model.Receipt) int {
	points := 0

	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			// convert price to float, perform mathematical operations on float, then round up to nearest int
			floatPrice := float64(item.Price.Cents)
			result := floatPrice / 100.0 * .2
			points += int(math.Ceil(result))
		}
	}
	return points
}

// OddDayRule 6 points if the day in the purchase date is odd.
type OddDayRule struct{}

func (rule *OddDayRule) evaluate(receipt *model.Receipt) int {
	points := 0
	if receipt.PurchaseDate.Day()%2 == 1 {
		points = 6
	}
	return points
}

// TwoPMTo4PMRule 10 points if the time of purchase is after 2:00pm and before 4:00pm.
type TwoPMTo4PMRule struct{}

func (rule *TwoPMTo4PMRule) evaluate(receipt *model.Receipt) int {
	points := 0

	// convert hh:mm to continuous integer values and compare
	lower := 14 * 60
	upper := 16 * 60
	time := receipt.PurchaseTime.Hour()*60 + receipt.PurchaseTime.Minute()
	// ASSUMPTION 2:00pm is not 'after' 2:00PM and 4:00pm is not 'before' 4:00PM
	// Evaluates to true if 2:01PM <= PurchaseTime <= 3:59PM
	if time > lower && time < upper {
		points = 10
	}
	return points
}
