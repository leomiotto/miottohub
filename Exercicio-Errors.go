package main

import (
	"fmt"
	"math"
)
/*func Sqrt(x float64) (float64, error) {
	return 0, nil
}*/

type NegSqrt float64

func (e NegSqrt) Error() string{
	return fmt.Sprintf("Negative number: %v", float64(e))
}

func Raiz(x float64) (float64, error){
	if x<0 {
		return 0, NegSqrt(x)
	}else {
		return math.Sqrt(x), nil
	}
}

func main() {
	fmt.Println(Raiz(9))
	fmt.Println(Raiz(-2))
}
