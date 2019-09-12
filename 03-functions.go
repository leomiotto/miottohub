package main

import "fmt"

func add (x, y int) int{
	return x + y
}

func main(){
	fmt.Println("The sum is: ", add(1,2))
}
