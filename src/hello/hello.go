package main

import (
	"fmt"
)

func main() {
	nome := "Carol"
	versao := 1.1

	fmt.Println("Bem vindo ao programa, ", nome)
	fmt.Println("Este programa está na versão: ", versao)

	//fmt.Println("O tipo da variavel é: ", reflect.TypeOf(nome))
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	fmt.Println("Endereço na memória da var comando: ", &comando)
	fmt.Println("O comando escolhido foi: ", comando)

}
