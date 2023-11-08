package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hi there Go!"
	fmt.Println(s1)

	fmt.Println(`hello ""gus"`)
	var1 := 'a'

	fmt.Println(s1[3:6])
	// string slicing

	// var1 is a rune literal, it
	// contains one character and
	// its in single string
	_ = var1

	// string package
	log := fmt.Println
	result := strings.Contains("kk", "op")

	strings.Count("cheese", "e") // returns 3 no of e

	_ = strings.ToUpper("uu")
	_ = strings.ToLower("huftf")
	_ = strings.EqualFold("go", "Go")
	_ = strings.Repeat("kk", 78)
	_ = strings.Replace("192.168.0.0", ".", ":", 2) // 2 is the number of values to replace
	//-1 replaces all occurrences

	_ = strings.Split("a,b,c", ",")
	_ = strings.Join([]string{"i", "learn", "go", "lang"}, " ")
	_ = strings.TrimSpace("    ktmkt togkokg    tkoko     ")
	_ = strings.Trim("...hello bpy !!!", ".!")

	log(result)
}
