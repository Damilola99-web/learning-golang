package main

import (
	"fmt"
	"time"
)

func main() {
	language := "python"
	switch language {
	case "golang":
		fmt.Println("You sabi go lang")

	case "python":
		fmt.Println("You sabi python")

	case "typescript", "javascript":
		fmt.Println("You sabi typescript")
	default:
		fmt.Println("You no sabi anything for this tech")	
	}


	date := time.Now()
	fmt.Println(date.Clock())
}
