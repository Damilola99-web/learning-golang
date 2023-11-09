package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type names []string

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (n names) print() {
	for _, v := range n {
		fmt.Println(v)
	}
}

func main() {
	const day = 24 * time.Hour
	fmt.Println(day)

	seconds := day.Seconds()
	fmt.Println(seconds)

	friends := names{"korede", "dammy", "noheem", "ibrahim"}
	friends.print()

	fmt.Println(strings.Repeat("=", 50))

	names.print(friends)

	var n int64 = 478675695685
	fmt.Println(time.Duration(n))

	rect1 := rectangle{width: 5, height: 7}
	fmt.Println(rect1.area())

	// empty interface allows any type
	type emptyInterface interface{}
	var rasheed emptyInterface = "ukorr"
	rasheed = 600
	fmt.Println(rasheed)


}
