package main

import "fmt"

func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	c <- f
}

func main() {
	channel := make(chan int)
	defer close(channel)
	go factorial(5, channel)
	f := <-channel
	fmt.Println(f)

	for i := 0; i <= 20; i++ {
		go factorial(i, channel)
		f := <- channel
		fmt.Println(f)
	}
}
