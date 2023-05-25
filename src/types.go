package main

import (
	"strconv"
	"strings"
	"time"
)

type Receipt struct {
	Id           string       `json:"id"`
	Retailer     string       `json:"retailer"`
	PurchaseDate Date         `json:"purchaseDate"`
	PurchaseTime Time         `json:"purchaseTime"`
	Total        PriceInCents `json:"total"`
	Items        []Item       `json:"items"`
	points       *int
}

type Item struct {
	ShortDescription string       `json:"shortDescription"`
	Price            PriceInCents `json:"price"`
}

type ProcessReceiptReturn struct {
	Id string `json:"id"`
}

type GetPointsReturn struct {
	Points int `json:"points"`
}

type PriceInCents struct {
	int
}

type Date struct {
	time.Time
}

type Time struct {
	time.Time
}

// UnmarshalJSON
//
//	Converts a price string, as received in the receipt, to an int representing number of cents
//	The int will be easier to work with when checking for equality and divisibility/**
func (price *PriceInCents) UnmarshalJSON(data []byte) error {
	str := bytesToString(data)
	float, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return err
	}

	*price = PriceInCents{int(float * 100)}
	return nil
}

func (date *Date) UnmarshalJSON(data []byte) error {
	str := bytesToString(data)
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}

	*date = Date{t}
	return nil
}

func (purchaseTime *Time) UnmarshalJSON(data []byte) error {
	str := bytesToString(data)
	t, err := time.Parse("15:04", str)
	if err != nil {
		return err
	}

	*purchaseTime = Time{t}
	return nil
}

func bytesToString(data []byte) string {
	return strings.Replace(string(data), "\"", "", -1)
}
