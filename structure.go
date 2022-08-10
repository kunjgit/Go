package main

import "fmt"

type student struct{
	name string
	age int 
	id int 


}

func main(){

	var kunj student
	kunj.age=18
	kunj.name="kunj"
	kunj.id=87
	fmt.Printf("the name of the struct variable is %v\n",kunj.name)
	fmt.Printf("the age of the struct variable is %d\n",kunj.age)
	fmt.Printf("the id of the struct variable is %d\n",kunj.id)

}