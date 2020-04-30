package main

import "fmt"

type Topping struct {
	name  string
	price float32
}

func NewTopping(name string, price float32) *Topping {
	return &Topping{
		name:  name,
		price: price,
	}
}

func (t *Topping) getItems() {
	fmt.Println(t)
}
