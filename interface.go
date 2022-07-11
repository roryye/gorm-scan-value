package main

import "errors"

type Animal interface {
	Eat(food string) error
}

func (a *AddressList) Eat(food string) error {
	if food != "" {
		return nil
	}
	return errors.New("food is empty")
}
