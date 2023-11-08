package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bs := make([]byte, 2)

	noBytesRead, err := io.ReadFull(file, bs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("no of bytes read", noBytesRead)
	fmt.Printf("%s\n", bs)

	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	data, err = ioutil.ReadFile("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data read: %s\n", data)
}
