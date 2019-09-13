//Map
package main

import "fmt"

type Vertex struct{
	Lat, Long float64
}

var m map[string]Vertex

func main(){
	m = make(map[string]Vertex)
	m["Joinville"] = Vertex{-26.3051, -48.8461}
	
	fmt.Println("Localização:",m["Joinville"])
}


//Map literal
package main

import "fmt"

type Vertex struct{
	Lat, Long float64
}

var m = map[string]Vertex{
	"Joinville": Vertex{-26.3051, -48.8461},   //Tipo "Vertex" aqui e na próxima linha facultativos, 
	"Floripa": Vertex{-27.5969,-48.5495},      //pois o enunciado tem o mesmo tipo.
}


func main(){
	fmt.Println("Localização:",m["Joinville"], m["Floripa"])
}

