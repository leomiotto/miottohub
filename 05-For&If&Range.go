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

	
	//Range
	pow := make([]int, 10)
	for i := range pow {   //Sem usar o value
		pow[i] = 2 << uint(i)   //O "<<" calcula o dobro de 2, 4, 8... i
	}
	for _, value := range pow {   //Sem usar o index
		fmt.Printf("%d\n", value)
	}
}
