package main

import "fmt" //Importa o Print

//Soma
func add (x int, y int) int{
	return x + y
}

//Concatena duas Strings - Duas variáveis do mesmo tipo
func concatena(x, y string) string{
	return x + y
}

//Separa dois números - Retorno especificado no enunciado (Naked Return)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main(){
	fmt.Println("The sum is: ", add(1,2))
	a := concatena("Olá,", " Marilene!")
	fmt.Println("A concatenação é:", a)
	fmt.Println("Números separados: ", split(15))
}
