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

	/*
		Tipos de inteiros:
			int8:   -128 a 127
			int16:  -32,768 a 32,767
			int32:  -2,147,483,648 a 2,147,483,647
			int64:  -9,223,372,036,854,775,808 a 9,223,372,036,854,775,807
			Sem sinal:
				uint8:  0 a 255
				uint16: 0 a 65,535
				uint32: 0 a 4,294,967,295
				uint64: 0 a 18,446,744,073,709,551,615
			Obs: Para inteiros sem tamanho especificado, se utiliza int ou uint, que vão ser definidos pela arquitetura do sistema (32 ou 64 bits)
		Tipos de float:
			float32: de -3.4E+38 a 3.4E+38 (6-7 dígitos de precisão)
			float64: de -1.7E+308 a 1.7E+308 (15-16 dígitos de precisão)
	*/

	var idade uint8 = 16
	var ano uint16 = 2024
	var populacao_itabaiana uint32 = 105_000 // _ como separador de milhar, facilita a leitura
	fmt.Printf("Idade (8 bits): %d anos; Ano (16 bits): %d; População de Itabaiana (32 bits): %d\n", idade, ano, populacao_itabaiana)

	var saldo_banco float32 = 1234.56
	fmt.Printf("Saldo bancário (float32): %.2f\n", saldo_banco)

	fmt.Print("Digite seu nome: ")
	var nome string
	// Pega a entrada do usuário e armazena na variável nome
	fmt.Scanf("%s", &nome) // Scanf requer o endereço da variável/ponteiro (onde a variavel será/está armazenada)
	fmt.Printf("Seu nome é %s? Nome feio esse..\n", nome)
}