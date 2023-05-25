package rules

import (
	"github.com/go-playground/assert/v2"
	"receipt-processor-challeng/src/model"
	"testing"
	"time"
)

func TestReatilerLengthRule(t *testing.T) {
	r := &model.Receipt{
		Retailer: "T3$t R3t@a!13r",
	}

	rule := &RetailerLengthRule{}
	points := rule.evaluate(r)
	assert.Equal(t, points, 10)
}

func TestRoundDollarAmountRule(t *testing.T) {
	r := &model.Receipt{
		Total: model.PriceInCents{100},
	}

	rule := &RoundDollarAmountRule{}
	points := rule.evaluate(r)
	assert.Equal(t, points, 50)

	r.Total = model.PriceInCents{101}
	points = rule.evaluate(r)
	assert.Equal(t, points, 0)
}

func TestPoint25Rule(t *testing.T) {
	r := &model.Receipt{
		Total: model.PriceInCents{125},
	}

	rule := &PointTwoFiveRule{}
	points := rule.evaluate(r)
	assert.Equal(t, points, 25)

	r.Total = model.PriceInCents{126}
	points = rule.evaluate(r)
	assert.Equal(t, points, 0)
}

func TestItemsLengthRule(t *testing.T) {
	r := &model.Receipt{
		Items: []model.Item{
			{},
			{},
			{},
		},
	}

	rule := &ItemsLengthRule{}
	points := rule.evaluate(r)
	assert.Equal(t, points, 5)

	r.Items = append(r.Items, model.Item{})
	points = rule.evaluate(r)
	assert.Equal(t, points, 10)
}

func TestTrimmedItemDescription(t *testing.T) {
	r := &model.Receipt{
		Items: []model.Item{
			{
				"Mountain Dew 12PK",
				model.PriceInCents{649},
			},
			{
				"Emils Cheese Pizza",
				model.PriceInCents{1225},
			},
			{
				"Knorr Creamy Chicken",
				model.PriceInCents{126},
			},
			{
				"Doritos Nacho Cheese",
				model.PriceInCents{335},
			},
			{
				"   Klarbrunn 12-PK 12 FL OZ  ",
				model.PriceInCents{1200},
			},
		},
	}

	rule := &TrimmedItemDescriptionRule{}
	points := rule.evaluate(r)
	assert.Equal(t, points, 6)
}

func TestOddDayRule(t *testing.T) {
	r := &model.Receipt{
		PurchaseDate: model.Date{time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	rule := &OddDayRule{}
	points := rule.evaluate(r)
	assert.Equal(t, points, 6)

	r.PurchaseDate = model.Date{time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)}
	points = rule.evaluate(r)
	assert.Equal(t, points, 0)
}

func TestAfternoonRule(t *testing.T) {
	r := &model.Receipt{
		PurchaseTime: model.Time{time.Date(2023, 1, 1, 14, 1, 0, 0, time.UTC)},
	}

	rule := &TwoPMTo4PMRule{}
	points := rule.evaluate(r)
	assert.Equal(t, points, 10)

	r.PurchaseTime = model.Time{time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC)}
	points = rule.evaluate(r)
	assert.Equal(t, points, 0)

	r.PurchaseTime = model.Time{time.Date(2023, 1, 1, 16, 0, 0, 0, time.UTC)}
	points = rule.evaluate(r)
	assert.Equal(t, points, 0)
}
