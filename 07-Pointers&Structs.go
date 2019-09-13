package main

import "fmt"

type Vertex struct {
	X, Y, Z int
}

//Formas de inicializar struct
var (
	v1 = Vertex{1,2,3}	 // has type Vertex
	v2 = Vertex{X:1, Y:3}	 // Z:0 is implicit
	v3 = Vertex{}     // X:0 Y:0 Z:0
	p = &Vertex{4,5,6}	 // has type *Vertex
)

func main() {
	i, j := 42, 2701
	//var p *int
	
	p := &i	// point to i
	fmt.Println(*p)	// read i through the pointer
	*p = 69	// set i through the pointer
	fmt.Println(i)	// see the new value of i

	p = &j	// point to j
	*p = *p/331	// divide j through the pointer
	fmt.Println(j)	// see the new value of j
  
	v := Vertex{1, 2, 3}
  p = &v
	p.X = 4   //Usar ponto pra acessa o campo da struct, o uso de (*p).X é facultativo
	fmt.Println(v.X)   //Print Struct
}
