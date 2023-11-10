package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println(url, " is DOWN")
	} else {
		defer res.Body.Close()
		fmt.Printf("%s -> Status Code: %d \n", url, res.StatusCode)
		if res.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			_ = err
			file := strings.Split(url, "//")[1]
			file += ".html"
			fmt.Println("writing response body to ", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Done()
}

func main() {
	urls := []string{"https://wahab-rasheed.vercel.app", "https://www.google.com", "https://unavailable.com", "https://golang.org"}
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
	}

	wg.Wait()
}
