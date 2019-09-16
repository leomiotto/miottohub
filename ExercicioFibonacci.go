package main

import "fmt"
//o próximo é o atual mais o passado
// fibonacci is a function that returns
// a function that returns an int.   fnext i = fcurrent i + fcurrent i - 1
func fibonacci() func() int {
	first, second := 0, 1
	return func() int{
		ret := first
		first = second
		second = first + second
		
		return ret
	}
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
