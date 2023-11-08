package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	text := scanner.Text()
	byte := scanner.Bytes()

	fmt.Println("input in text : ", text)
	fmt.Println("input in byte : ", byte)
}