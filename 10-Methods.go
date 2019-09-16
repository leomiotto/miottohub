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


//POINTER RECEIVER - Usado pra modificar o valor do argumento Receiver
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


//INTERFACE
package main

import (
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perim() float64
}

type rect struct{
	width, height float64
}

type circle struct{
	radius float64
}

//METHODS
func (r rect) area() float64{   //float64 é o tipo de retorno
	return r.width * r.height   //retorna a fórmula da área
}
func (c circle) area() float64{
	return c.radius * math.Pi * c.radius
}

func (r rect) perim() float64{
	return 2*r.width + 2*r.height
}

func (c circle) perim() float64{
	return 2 * math.Pi * c.radius
}

//INTERFACE
func measure (g geometry){
	fmt.Println("Medidas:", g)   //ex: {3,4}
	fmt.Println("Área: ", g.area())
	fmt.Println("Perímetro: ", g.perim())
}

func main(){
	r1 := rect{width: 3, height: 4}
	c1 := circle{radius: 5}
	//rv := &r1   -   Para passar Pointer Receiver
	measure(r1)
	measure(c1)
}
