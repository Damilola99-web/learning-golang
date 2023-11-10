package main

import (
	"fmt"
	"sync"
)

func main() {
	const gr = 100

	var n int = 0

	var m sync.Mutex

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()
		go func() {
			m.Lock()
			defer m.Unlock()
			n--
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n)
}
