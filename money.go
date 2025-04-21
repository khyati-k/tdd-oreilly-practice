package main

type Money struct {
	amount float64
	currency string
}

func NewMoney(amount float64, currency string) Money {
	if amount < 0 {
        panic("amount cannot be negative")
    }
	return Money{amount: amount, currency: currency}
}

func (m Money) Times(factor float64) Money {
	return Money{amount: m.amount * factor, currency: m.currency}
}

func (m Money) Divide(factor float64) Money {
	return NewMoney(m.amount / factor, m.currency)
}
