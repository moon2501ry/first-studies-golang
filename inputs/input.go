package main

import "fmt"

func main() {
	fmt.Print("Digite sua altura e peso separados por espaço: ")
	var altura float32
	var peso float32
    // Pega a entrada do usuário e armazena na variável
	fmt.Scanf("%f %f", &altura, &peso) // Scanf requer o endereço da variável/ponteiro (onde a variavel será/está armazenada)
	fmt.Printf("Sua altura é %.2f metros e seu peso é %.2f kg\n", altura, peso)
	fmt.Printf("Seu IMC é %f\n", peso/(altura*altura))
}
