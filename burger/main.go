package main

func main() {
	r := NewRestaurant("THE BURGER SHOP")

	r.displayMenu()
	r.selectBurger()
	r.displayToppings()
	r.selectToppings()
	r.generateBill()
}
