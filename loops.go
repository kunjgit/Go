package main

import (
	"fmt"
)

func main(){
 
 for i:=0;i<10;i++{
	fmt.Println("the value of this number is just",i)

 }


 	numbers :=[7]int {1,2,3,44,5}

	 for i,x:= range numbers {
		fmt.Printf("value of x = %d at %d\n", x,i)
	 }   


	 fmt.Println("another syntax to print the array in go")
	 fmt.Println(numbers)


// var a="hello world"
// var age=23  
	// fmt.Println(kunj)
}