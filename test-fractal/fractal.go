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

	fmt.Print("Valor de redução: ")
	var reduc float32
	fmt.Scanf("%f\n", &reduc)

	var max_v float64 = 0
	var max_h float64 = 0
	for i := 1; i <= iteracoes; i++ {
		max_v += math.Abs(math.Cos(float64(float32(i)*angulo)/2)) * float64(a/(reduc*float32(i)))
		max_h += math.Abs(math.Sin(float64(float32(i)*angulo)/2)) * float64(a/(reduc*float32(i)))
	}
	fmt.Printf("Crescimento Máximo Vertical: %f\nCrescimento Máximo Horizontal: %f\nAltura Máxima: %f\nLagura Máxima: %f\n", max_v, max_h, max_v+float64(a), max_h*2)
}