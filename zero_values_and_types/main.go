package main

import "fmt"

func main() {
	var a = 5
	var b = 4.8

	a = int(b)

	fmt.Println(a, b)
}

// zero values, no variable is undeclared in golang
// numeric 0
// bool false
// string ""
// pointer nil
