package main

import (
	"fmt"
	"math/rand"
)

func main(){
	rand.Seed(88)
	fmt.Println("Número", rand.Intn(152))
}
