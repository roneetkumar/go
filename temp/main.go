package main

import "fmt"

func main() {

	// container := [][]int32{
	// 	{0, 2, 1},
	// 	{1, 1, 1},
	// 	{2, 0, 0},
	// }

	container := [][]int32{
		{1, 3, 1},
		{2, 1, 2},
		{3, 3, 3},
	}

	organizingContainers(container)

}

// Complete the organizingContainers function below.
func organizingContainers(container [][]int32) {

	fmt.Println(container)

	for c := 0; c < len(container); c++ {
		for t := 0; t < len(container); t++ {
			if 0 == t {
				continue
			}
			container[0][c] += container[t][c]
		}
	}

	possible := "Possible"

	for _, i := range container[0] {
		if int(i) > len(container) {
			possible = "Impossible"
		}
	}

	fmt.Println(possible, container)

	// return possible
}
