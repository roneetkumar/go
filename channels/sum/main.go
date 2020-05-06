package main

import "fmt"

func main() {

	fmt.Println("Add numbers with Goroutines")

	a := []int{2, 3, 4, 5, 6, 1, 54, 7, 3, 6, 3, 4, 5}

	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}

func sum(a []int, c chan int) {
	total := 0
	for _, n := range a {
		total += n
	}
	c <- total
}
