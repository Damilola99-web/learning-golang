package main

import "fmt"

func main() {
	var channel chan int 
	channel = make(chan int)

	// or
	channel2 := make(chan int)

	// send operation 
	channel <- 50

	// recieve expression 
	num := <- channel2
	fmt.Println(num)
	fmt.Println(<- channel)

	// closing channel 
	close(channel)
	
}