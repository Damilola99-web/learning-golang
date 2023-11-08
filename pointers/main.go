package main

import "fmt"

func main() {
	x := 5
	pointer := &x
	fmt.Println(pointer)

	p := new(int)
	x = 100
	p = &x

	*p = 10 //x =10

	fmt.Println(x)

	// &value => pointer
	// *pointer => value

	a := 5.5
	p1 := &a
	pp1 := &p1

	**pp1 = 6.6
	fmt.Println(a)

}
