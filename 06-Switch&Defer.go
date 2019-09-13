package main

import (
	"fmt"
	"math/rand"
)

func main(){
	rand.Seed(77)
	var a int = rand.Intn(431)
	
	switch {   //Posso colocar condição aqui ou não, Switch without a condition is the same as switch true.
		case a < 150:
			fmt.Println("Pouca coisa ", a)
		case a < 300:
			fmt.Println("Até que sim ", a)
		default:
			fmt.Println("Baita ", a)
	}
  
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)   //Espera um próximo retorno para retornar após ele.
	}

	fmt.Println("done")
}
