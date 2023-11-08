package main

import "fmt"

func main() {
	// array
	accounts := [3]int{20, 30, 40}

	var numbers [4]int

	var a1 = [4]float64{}

	// elipsis operator
	names := [...]string{"rasheed", "korede"}

	// modifying array elements
	names[0] = "bululu"

	fmt.Println(numbers, a1, names)

	// iterating over array
	for i, v := range names {
		fmt.Println("index:", i, " value:", v)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println("I am ", names[i])
	}

	_ = accounts
}
