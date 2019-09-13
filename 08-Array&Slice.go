package main

import (
    "fmt"
    "strings"   //Linha 76
)

func main() {
    //Arrays
	var a [2]string
	a[0] = "Hello"   //Arrays cannot be resized.
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
    
    
    //Slices
	primes1 := [6]int{2, 3, 5, 7, 11, 13}   //A slice is a dynamically-sized, flexible view into the elements of an array.
	var s []int = primes1[1:6]   //Slice não armazena dados, só referencia um vetor
	fmt.Println(s)
    
    names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)   //[John Paul George Ringo]

	a1 := names[0:2]   //Slice pega a posição 0 mas não a posição 2 (sempre inclui o primeiro e exclui o último)
	b := names[1:3]
	fmt.Println(a1, b)   //[John Paul] [Paul George]

	b[0] = "XXX"
	fmt.Println(a1, b)   //[John XXX] [XXX George]
	fmt.Println(names)   //[John XXX George Ringo]
    
    //Slice Literal; ele cria um array(tamanho fixo) já o referenciando
    s1 := []struct {
		i int
		b bool
	}{   //Vai criar um array (var s [6]struct)
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s1)
    
    //Slice com a função make; cria array com zeros e slice referenciando
    b2 := make([]int, 0, 5)   // len(b)=0, cap(b)=5

    b2 = b2[:cap(b)]   // len(b)=5, cap(b)=5
    b2 = b2[1:]      // len(b)=4, cap(b)=4
    
    
    // Create a tic-tac-toe board.
	board := [][]string{   //Slices of slices
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
    
    //Função para adicionar elementos ao Slice
    //func append(s []T, ele1,..., elen T) []T   junta ao slice "s" de tipo "array de T" elementos do tipo "T"
}
