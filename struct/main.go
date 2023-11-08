package main

import "fmt"

func main() {
	title1, author1, year1 := "Atomic habits", "James Clear", 2002
	title2, author2, year2 := "The Alchemist", "Unknown", 2016

	log := fmt.Println

	type book struct {
		title  string
		author string
		year   int
	}
	book1 := book{title1, author1, year1}
	book2 := book{title: title2, author: author2, year: year2}

	book1.title = "kokoko"

	log(book1, book2.title)

	// anonymous structs
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       40,
	}
	log(diana)

	type anonymous struct {
		string
		float64
		bool
	}
	anon := anonymous{"uhur", 10.2, false}
	log(anon.string)

	// embeded struct 
	type Contact struct {
		email, address string
		phone int
	}
	type Employee struct {
		name string
		salary int 
		contactInfo Contact
	}
}
