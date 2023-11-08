package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	success := scanner.Scan()
	if success == false {
		err = scanner.Err()
		if err == nil {
			log.Printf("Scan was completed")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("first line : ", scanner.Text())

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
