package main

import (
	"fmt"
	"strconv"
)

type Restaurant struct {
	name     string
	fridge   *Fridge
	burger   *Burger
	toppings []*Topping
}

// NewRestaurant func
func NewRestaurant(name string) *Restaurant {
	return &Restaurant{
		name:   name,
		fridge: NewFridge(),
	}
}

func (r Restaurant) displayMenu() {
	fmt.Printf("%s\n", r.name)
	fmt.Println("```````````````")
	for i, b := range r.fridge.burgers {
		fmt.Printf("%v. %v - $%v\n", i+1, b.name, b.price)
	}
}

func (r *Restaurant) selectBurger() {
	choice := selectItem()
	fmt.Println("select a burger")
	if choice > 3 || choice < 1 {
		fmt.Println("Invalid choice")
		r.selectBurger()
	} else {
		r.burger = r.fridge.burgers[choice-1]
	}
}

func (r Restaurant) displayToppings() {
	fmt.Println("\nAvailable Toppings")
	for i, t := range r.burger.toppings {
		fmt.Printf("%v. %v - $%v\n", i+1, t.name, t.price)
	}
}

func (r *Restaurant) selectToppings() {
	i := len(r.burger.toppings)
	for {
		fmt.Printf("choose any %d toppings : ", i)
		i--
		choice := selectItem()

		if choice == 0 {
			return
		} else if choice > len(r.burger.toppings) {
			fmt.Println("Please select the right topping")
			r.selectToppings()
		} else {
			r.toppings = append(r.toppings, r.burger.toppings[choice-1])
		}

		if len(r.burger.toppings) == len(r.toppings) {
			break
		}
	}
}

func selectItem() int {
	var choice string
	fmt.Scan(&choice)
	pchoice, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Println("select error")
	}
	return pchoice
}

func (r *Restaurant) generateBill() {
	amount := r.burger.price + toppingsTotal(r.toppings)
	tax := amount * 0.15
	total := amount + tax
	fmt.Println("Reciept")
	fmt.Println("```````")
	fmt.Printf("1. %v - %v\n", r.burger.name, r.burger.price)

	if len(r.toppings) > 0 {
		fmt.Println("Extras : ")
		fmt.Println("`````````")

		for i, top := range r.toppings {
			fmt.Printf("%v. %v - %v\n", (i + 1), top.name, top.price)
		}
	} else {
		fmt.Println("No Extra Toppings")
	}

	fmt.Printf("\nAmount: $%v", amount)
	fmt.Printf(" + (Tax: $%v)\n", tax)
	fmt.Printf("Total: $%v\n", total)
}

func toppingsTotal(arr []*Topping) float32 {
	var total float32
	for _, item := range arr {
		total += item.price
	}
	return total
}
