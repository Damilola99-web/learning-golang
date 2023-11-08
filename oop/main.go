package main

import (
	"fmt"
	"strings"
	"time"
)

type names []string

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

}
