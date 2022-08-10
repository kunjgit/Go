package main

import "fmt"

func fact(n int) int{
	if(n==0){
		return 1
	}

	return n*fact(n-1)

}


func main (){
//	var a int
//	a=4
	var ans=fact(10);

	fmt.Printf("the factorial of number %d is %d ", 10,ans)

}
