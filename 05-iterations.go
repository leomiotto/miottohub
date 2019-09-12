package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 10; i++ {  //Primeiro e último argumentos são opcionais, caso não haja vira "while"
		sum += i
		if v := 5; i >= v {  //Posso inicializar variavél antes da condição. Obs: ela só vale para blocos "if" e "else".
			fmt.Println("i está maior que v")
		} else {
			fmt.Println("v está maior que i")
		}
	}

	fmt.Println(sum)
}
