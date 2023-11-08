package main

import (
	"fmt"
)

func main() {

	employees := map[string]string{}
	fmt.Println(len(employees))
	_ = employees

	// accessing key in map
	_ = employees["hr"]

	// inserting key value pair
	employees["lead"] = "rasheed"

	people := map[string]string{
		"NGA": "nigeria",
		"USA": "usa",
		"BEL": "belgium",
	}

	// if the key exists, it updates it value

	_ = people

	value, ok := people["FIA"]
	log := fmt.Println
	if ok {
		log(value)
	} else {
		log("the key dosent exist")
	}

	// deleting key pair
	delete(people, "BEL")

	// looping through map
	for key, value := range people {
		log(key, value)
	}
}
