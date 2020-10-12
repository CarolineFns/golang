package main

import (
	"fmt"
)

func main() {
	nome := "Carol"
	versao := 1.1

	fmt.Println("Olá,", nome, "!")
	fmt.Println("Bem vindo ao programa de monitoramento")
	fmt.Println("Este programa está na versão:", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Println("Digite o comando desejado:")
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)

	switch comando {
	case 1:
		fmt.Println("Monitoramento...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Comando não reconhecido.")
	}

}
