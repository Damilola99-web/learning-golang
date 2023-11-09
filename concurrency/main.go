package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 start")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1 exit")
	wg.Done()
}
func f2() {
	fmt.Println("f2 start")
	for i := 5; i < 8; i++ {
		fmt.Println(i)
	}
	fmt.Println("f2 exit")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println("no of cpus: ", runtime.NumCPU())
	fmt.Println("no of goroutine: ", runtime.NumGoroutine())

	fmt.Println("os", runtime.GOOS)
	fmt.Println("architecture", runtime.GOARCH)
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))

	// calling a go routine
	go f1(&wg)
	fmt.Println("no of goroutine: ", runtime.NumGoroutine())
	f2()
	// time.Sleep(time.Second * 2)
	wg.Wait()
	fmt.Println("main execution stopped")
}
