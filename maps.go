package main

import "fmt"

func main(){
	var a =make(map[string]string) 
	a["kunj"]="hello"
	a["kundsajh"]="hesfdl"
	a["kuasfklj"]="heasfkjh"
	a["safkljlsnj"]="sfjnkahof"

	for a,b:=range a{
		fmt.Printf("%v at %v\n",a,b)
	}

}