package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Valor de A (tamanho do primeiro galho): ")
	var a float32
	fmt.Scanf("%f\n", &a)

	fmt.Print("Valor de B (ângulo da primeira ramificação): ")
	var angulo float32
	fmt.Scanf("%f\n", &angulo)

	fmt.Print("Valor de i (iterações): ")
	var iteracoes int
	fmt.Scanf("%d\n", &iteracoes)

	for i := 1; i >= iteracoes; i++ {
		math.Cos((float32(i) * angulo))
	}
}
