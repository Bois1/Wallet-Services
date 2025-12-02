package main

import (
	"errors"
	"fmt"
)


type Money int64


func NewMoney(kobo int64) (Money, error) {
	if kobo <= 0 {
		return 0, errors.New("amount must be a positive number")
	}
	return Money(kobo), nil
}


func (m Money) Kobo() int64 {
	return int64(m)
}


func (m Money) Naira() int64 {
	return m.Kobo() 
}


func (m Money) String() string {
	naira := m.Naira()
	kobo := m.Kobo() % 100
	return fmt.Sprintf("₦%d.%02d", naira, kobo)
}