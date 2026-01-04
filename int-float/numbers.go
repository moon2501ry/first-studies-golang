package main

import "fmt"

func main() {
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
}
