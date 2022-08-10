package main

import "fmt"

func main() {
	sl := []int{2, 3, 45, 6, 6, 6, 7}
	sl1 := make([]string, 4, 5)
	sl2 := []string{"slices", "are", "more", "powerful than array"}
	fmt.Printf("the information for sl capacity %d and size %d\n", len(sl), cap(sl))
	fmt.Printf("the information for sl capacity %d and size %d\n", len(sl1), cap(sl1))
	fmt.Printf("the information for sl capacity %v\n", sl2)

	myslice1 := []int{1,2,3}
	myslice2 := []int{4,5,6}
	myslice3 := append(myslice1, myslice2...)
  
	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))

}
