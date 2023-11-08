package main

import "fmt"

func main() {
	const pie float64 = 3.142
	const n, m int = 4, 5

	const (
		min  int    = -500
		max  int    = 500
		rate string = "dollar"
	)

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3) //0 1 2

	const (
		c11 = iota
		_
		c22 = iota
		c33 = iota
	)

	fmt.Println(c11, c22, c33) //0 2 3

	fmt.Printf("Max : %d, Min : %d, Conversion Rate : %s \n", min, max, rate)
}
