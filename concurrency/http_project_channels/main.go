package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrl(url string, ch chan string) {
	res, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf(url, " is DOWN")
		s += fmt.Sprintf("Error: %v\n", err)
		fmt.Println(s)
		ch <- url
	} else {
		s := fmt.Sprintf("%s -> Status Code: %d \n", url, res.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)
		fmt.Println(s)
		ch <- url

	}
}

func main() {
	urls := []string{"https://wahab-rasheed.vercel.app", "https://www.google.com", "https://golang.org"}

	ch := make(chan string)

	for _, url := range urls {
		go checkUrl(url, ch)

	}
	// for {
	// 	go checkUrl(<-ch, ch)
	// 	fmt.Println(strings.Repeat("*", 20))
	// 	time.Sleep(time.Second * 5)
	// }

	for url := range ch {
		time.Sleep(time.Second * 5)
		go checkUrl(url, ch)
	}
}
