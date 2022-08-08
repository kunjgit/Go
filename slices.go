package main

import "fmt"

func main() {
	sl := []int{2, 3, 45, 6, 6, 6, 7}
	sl1 := make([]string, 4, 5)
	sl2 := []string{"slices", "are", "more", "powerful than array"}
	fmt.Printf("the information for sl capacity %d and size %d\n", len(sl), cap(sl))
	fmt.Printf("the information for sl capacity %d and size %d\n", len(sl1), cap(sl1))
	fmt.Printf("the information for sl capacity %v\n", sl2)

}