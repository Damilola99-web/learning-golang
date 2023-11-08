package main

import "fmt"

func test() {
	fmt.Println("function")
}

// naked return
func f2(a int, b int) (z int) {
	z = a + b
	return
}

// variadic function
func add(a ...int) {
	fmt.Println(a)
}

func f3(a, b, c int, s string) {
	fmt.Println(a, b, c, s)
}

// defer statement
func foo() {
	defer test()
	fmt.Println("this is foo")
}

func main() {
	test()
	fmt.Println(f2(6, 5))
	f3(2, 3, 4, "iuyrg")
	add(2, 3, 45, 6)
	foo()
	// anonymous function
	func(msg string) {
		fmt.Println(msg)
	}("hello")
}
