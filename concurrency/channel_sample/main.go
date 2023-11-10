package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	channel := make(chan int)

	// only for receiving
	c1 := make(<-chan string)

	// only for sending
	c2 := make(chan<- string)

	fmt.Printf("%T,%T,%T \n", channel, c1, c2)

	go f1(10, channel)

	n := <-channel
	fmt.Println( n)
}
