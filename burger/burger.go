package main

import "fmt"

type Burger struct {
	name     string
	price    float32
	roll     string
	meat     string
	toppings []*Topping
}

func NewBurger(name string, price float32, roll string, meat string) *Burger {
	return &Burger{
		name:  name,
		price: price,
		roll:  roll,
		meat:  meat,
	}
}

func (b *Burger) addToppings(topping *Topping) {
	b.toppings = append(b.toppings, topping)
}

func (b *Burger) getItems() {
	fmt.Println(b)
}
