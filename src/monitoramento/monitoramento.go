package main

import (
	"fmt"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitoramento...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando não reconhecido.")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Carol"
	versao := 1.1

	fmt.Println("Olá,", nome, "!")
	fmt.Println("Bem vindo ao programa de monitoramento")
	fmt.Println("Este programa está na versão:", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Println("Digite o comando desejado:")
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
