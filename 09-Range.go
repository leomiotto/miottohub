package main

import "fmt"

func main(){
	pow := make([]int, 10)
	for i := range pow {   //Sem usar o value
		pow[i] = 2 << uint(i)   //O "<<" calcula o dobro de 2, 4, 8... i
	}
	for _, value := range pow {   //Sem usar o index
		fmt.Printf("%d\n", value)
	}
}
