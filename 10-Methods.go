package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X,Y float64
}

func (v Vertex) hip() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
	
}

func main(){
	v := Vertex{3,4}
	fmt.Println(v.hip())
}


//Pointer Receiver - Usado pra modificar o valor do argumento Receiver
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y) //hipotenusa
}

func (v *Vertex) Scale(f float64) { // Como é passado ponteiro ele consegue modificar o valor de v
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) //Method: O Go interpreta v.Scale(5) como (&v).Scale(5), já que a função deveria receber ponteiro
	fmt.Println(v.Abs())
}
