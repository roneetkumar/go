package main

type Fridge struct {
	burgers  []*Burger
	toppings []*Topping
}

func NewFridge() *Fridge {

	bb := NewBurger("Basic Burger", 4.2, "White", "Chicken")
	hb := NewBurger("Healthy Burger", 5.4, "White", "Chicken")
	db := NewBurger("Deluxe Burger", 8.9, "White", "Chicken")

	bb.addToppings(NewTopping("Carrot", 2.21))
	bb.addToppings(NewTopping("Onion", 3.29))
	bb.addToppings(NewTopping("Pepper", 1.25))
	bb.addToppings(NewTopping("Cheese", 3.42))

	hb.addToppings(NewTopping("Carrot", 2.21))
	hb.addToppings(NewTopping("Onion", 3.29))
	hb.addToppings(NewTopping("Pepper", 1.25))
	hb.addToppings(NewTopping("Cheese", 3.42))
	hb.addToppings(NewTopping("Lettuce", 2.32))
	hb.addToppings(NewTopping("Tamato", 1.12))

	db.addToppings(NewTopping("Lentils", 4.32))
	db.addToppings(NewTopping("Mayo", 1.52))
	var burgers []*Burger

	burgers = append(burgers, bb, hb, db)

	return &Fridge{
		burgers: burgers,
		// toppings,
	}
}
