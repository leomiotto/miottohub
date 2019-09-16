//METHODS
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
	"reflect"
)

type geometria interface{   //Interface
	area() float64
	perimetro() float64
}

type retangulo struct{
	altura, largura float64
}

type circulo struct{
	raio float64
}

func (r retangulo) area() float64{
	return r.altura * r.largura
}

func (c circulo) area() float64{
	return math.Pi * math.Pow(c.raio, 2)
}

func (r retangulo) perimetro() float64{
	return 2*r.altura + 2*r.largura
}

func (c circulo) perimetro() float64{
	return 2 * math.Pi * c.raio
}

func print_medidas (g geometria){
	var r retangulo
	var c circulo
	if reflect.TypeOf(g) == reflect.TypeOf(r){
		fmt.Println("Retângulo: ")
	}else if reflect.TypeOf(g) == reflect.TypeOf(c){
		fmt.Println("Círculo: ")
	}
	fmt.Println("Medidas: ", g)
	fmt.Println("Área: ", g.area())
	fmt.Println("Perímetro: ", g.perimetro())
	fmt.Println()
}

func main(){
	r1 := retangulo{altura: 2, largura: 4}
	r2 := retangulo{altura: 5, largura: 3}
	c1 := circulo{raio: 3}
	
	print_medidas(r1)
	print_medidas(c1)
	print_medidas(r2)
}
