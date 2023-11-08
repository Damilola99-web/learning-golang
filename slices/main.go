package main

import "fmt"

func main() {
	// declaring slice
	var cities []string
	fmt.Println(cities) //nil or []

	schools := []string{"unilorin", "unilag", "futa", "funnab"}
	fmt.Println(schools)

	nums := make([]int, 50)
	num := []int{3, 4, 5}
	fmt.Println(nums)

	for index, value := range schools {
		fmt.Println(index, value)
	}

	// appending to a slice
	schools = append(schools, "lasu", "ui")
	fmt.Println(schools)

	// joining two slices together
	numbers := append(nums, num...)
	fmt.Println(numbers)

	// copying slices
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(dst, src)

	fmt.Println(dst)

	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[1:3]
	s3 := s1[:3]
	s4 := s1[1:]
	s5 := s1[:]
	fmt.Println(s2, s3, s4,s5)

	// slices can only be compared to nil
}
