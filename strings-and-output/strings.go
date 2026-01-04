package main

import (
	"fmt"
	"strings"
)

func main() {
    var str1 string   // Declara um string vazia; var str1 string = ""
	str1 = "String1"  // Atribui um valor à ela
	str2 := "String2" // Declara e atribui um valor à ela

	fmt.Println("Println")                                // Imprime uma mensagem e pula uma linha
	fmt.Print("Print")                                    // Imprime uma mensagem sem pular uma linha
	fmt.Printf("\nPrint with format %s %s\n", str1, str2) // Imprime uma mensagem formatada

	/*
		Tipos de formatação comuns:
		%s - string
		%c - caractere (byte/ASCII)
		%d - inteiro decimal
		%f - ponto flutuante (float)
			Pode ser especificado o número de casas decimais: %.Xf (X casas)
		%t - booleano
		Obs: Para imprimir o caractere % utiliza-se %%
	*/
	alfabeto := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Printf("Alfabeto: %s\n", alfabeto)
	primeira_letra := alfabeto[0]                                // Acessa o caracter em formato de byte (ASCII)
	fmt.Printf("Primeira letra (index 0): %c\n", primeira_letra) // Formata o byte como um caractere

	tamanho_string := len(alfabeto)    // Tamanho da string
	ultimo_index := tamanho_string - 1 // Último índice da string
	ultima_letra := alfabeto[ultimo_index]
	fmt.Printf("Última letra (index %d): %c\n", ultimo_index, ultima_letra)

	abcd := alfabeto[0:4] // Fatiamento da string (slice), o ultimo índice é exclusivo (não incluso)
	fmt.Printf("Primeiras 4 letras do alfabeto: %s\n", abcd)
	fmt.Printf("Ultimas 4 letras do alfabeto: %s\n", alfabeto[22:]) // Não é necessário colocar o último índice da string

	frase := "O Lato Loeu a Loupa do Lei de Loma"
	fmt.Printf("Frase: %s\n", frase)
	frase_minusculo := strings.ToLower(frase) // Converte todos os caracteres para minúsculo
	fmt.Printf("Frase em minúsculo: %s\n", frase_minusculo)
	frase_maiusculo := strings.ToUpper(frase) // Converte todos os caracteres para maiúsculo
	fmt.Printf("Frase em maiúsculo: %s\n", frase_maiusculo)

	frase_suja := "   Frase "
	fmt.Printf("Frase suja (tamanho %d): '%s'\n", len(frase_suja), frase_suja)
	frase_limpa := strings.TrimSpace(frase_suja) // Remove espaços em branco, ou Trim para outros caracteres
	fmt.Printf("Frase limpa (tamanho %d): '%s'\n", len(frase_limpa), frase_limpa)
}
