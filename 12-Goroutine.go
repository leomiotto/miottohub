package main

import (
	"fmt"
	"time"
)

func manda(msg chan string, feito1 chan bool){
	fmt.Println("Mandando")
	for i:=0 ; i<10 ; i++{
		time.Sleep(500 * time.Millisecond)
		fmt.Printf(".")
	}
	msg <- "Olá, Marilene!"
}

func recebe(msg chan string, feito1, feito2 chan bool){
	fmt.Println("Oi2")
	<-feito1
	
	fmt.Println("Recebendo...")
	fmt.Println(<-msg)
	
	for i:=0 ; i<10 ; i++{
		time.Sleep(200 * time.Millisecond)
		fmt.Printf(".")
	}
	
	feito2<-true
}

func main() {
	msg := make(chan string)
	feito1 := make(chan bool)
	feito2 := make(chan bool)
	
	go manda(msg, feito1)
	go recebe(msg, feito1, feito2)
	
	fmt.Println("Olá, World!")

	<-feito2
}




/*package main

import (
	"fmt"
	"time"
)

func task(savior chan bool){
	for i:=0 ; i<10 ; i++{
		fmt.Println("START")
		fmt.Println(i, i, i, i, i, i, i, i, i, i)
		fmt.Println("END")
		time.Sleep(1 * time.Second)
	
		savior <- true
	}
}

func main() {
	canalis := make(chan bool)
	
	go task(canalis)
	fmt.Println("Olá, World!")
	
	for i:=0 ; i<10 ; i++{
		<- canalis
	}
}
*/


//Assinatura da função, limita a função
//(done chan<- bool) //Canal só recebe valor
//(done<- chan bool) //Canal só passa valor
