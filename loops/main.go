package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
	}

	// while loop
	j := 10
	for j >= 0 {
		// fmt.Println(j)
		j--
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
