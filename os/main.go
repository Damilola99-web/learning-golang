package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required")
	} else {
		if number, err := strconv.Atoi(args[1]); err == nil {
			fmt.Println("the number you passed is", number)
		} else {
			fmt.Println("You passed in an invalid argument")
		}
	}
}
