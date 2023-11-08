package main

import "fmt"

func main() {
	// specifying types
	var name string = "Rasheed"
	fmt.Println(name)

	var age int = 50
	fmt.Println(name, "is", age, "years old.")

	// without types i.e type inference
	var name2 = "basit"

	// you can use underscore to prevent error of unused vars
	_ = name2

	// short declaration operator
	name3 := "sheriff"
	fmt.Println(name3)

	// multiple declaration
	car, cost := "Benz", 6000
	fmt.Println(car, "costs", cost)

	var (
		nameee string = "korede"
		ageee  int    = 50
		gender string = "male"
	)
	_, _, _ = nameee, ageee, gender

	var bola, tolu, shayo int = 50, 60, 40
	_, _, _ = bola, tolu, shayo

}
