package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var newfile *os.File

	var err error

	newfile, err = os.Create("a.txt")

	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	} else {
		fmt.Println(newfile)
	}

	err = os.Truncate("a.txt", 0)
	if err != nil {
		log.Fatal(err)
	}

	newfile.Close()

	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(file)
	}

	var fileInfo os.FileInfo

	fmt.Println(strings.Repeat("=", 50))

	fileInfo, err = os.Stat("aba.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("name", fileInfo.Name())
		fmt.Println("size", fileInfo.Size())
		fmt.Println("last modifed", fileInfo.ModTime())
		fmt.Println("is directory?", fileInfo.IsDir())
		fmt.Println("permissions", fileInfo.Mode())
	}

	oldPath := "a.txt"
	newPath := "aba.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	
}
