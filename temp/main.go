package main

import "fmt"

func main() {

	a1 := []int32{2, 4}
	a2 := []int32{16, 32, 96}

	getTotalX(a1, a2)

}

func getTotalX(a []int32, b []int32) {

	temp := []int32{}
	for i := int32(0); i < 100; i++ {
		for _, n := range a {
			if i%n == 0 {
				temp = append(temp, i)
			}
		}

	}

	fmt.Println(temp)

}

func divisible(a int32) []int32 {
	temp := []int32{}
	for j := int32(1); j < 100; j++ {
		if j%a == 0 {
			temp = append(temp, j)
		}
	}
	return temp
}
