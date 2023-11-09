package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {
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
}

func main() {
	urls := []string{"https://wahab-rasheed.vercel.app", "https://www.google.com", "https://unavailable.com", "https://golang.org"}

	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Printf("\n\n")
	}
}
