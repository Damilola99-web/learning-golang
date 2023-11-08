package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	dob := "29-10-2005"

	presentYear := 2023
	presentMonth := 10
	presentDay := 30

	date := strings.Split(dob, "-")
	day, err := strconv.Atoi(date[0])
	_ = err
	month, err := strconv.Atoi(date[1])
	_ = err
	year, err := strconv.Atoi(date[2])
	_ = err

	if (month > 12) || month < 1 {
		fmt.Println("invalid month")
	}

	if year < 1000 {
		fmt.Println("Invalid year")
	}

	// thirties := []int{7, 4, 6, 11}

	if day > 31 || day < 1 {
		fmt.Println("invalid day")
	}

	// if thirties {

	// }

	if month > presentMonth {
		fmt.Println("Age is", presentYear-year-1)
	}

	if month == presentMonth && presentDay < day {
		fmt.Println("Age is", presentYear-year-1)
	}

	if month == presentMonth && presentDay >= day {
		fmt.Println("Age is", presentYear-year)
	}

	if presentMonth > month {
		fmt.Println("Age is", presentYear-year)
	} else if 2 == 3 {
		fmt.Println("hhh")
	} else {
		fmt.Println("yeey")
	}

}
