package main

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	tenner :=  NewMoney(10, "USD")
	result := tenner.Times(2)
	if result.amount != 20 {
		t.Errorf("Expected 20 but got %f", result.amount)
	}
}

func TestMultiplicationinEUR( t *testing.T) {
	tenner :=  NewMoney(10, "EUR")
	twentyEuros := tenner.Times(2)
	if twentyEuros.amount != 20 {
		t.Errorf("Expected 20 but got %f", twentyEuros.amount)
	}
	if twentyEuros.currency != "EUR" {
		t.Errorf("Expected EUR but got %s", twentyEuros.currency)
	}
}

func TestDivision(t *testing.T) {
	tenner :=  NewMoney(10, "USD")
	expectedMoneyAfterDivision := NewMoney(5, "USD")
	actualMoneyAfterDivision := tenner.Divide(2)
	if actualMoneyAfterDivision != expectedMoneyAfterDivision {
		t.Errorf("Expected [%+v] but got [%+v]", expectedMoneyAfterDivision, actualMoneyAfterDivision)
	}
}