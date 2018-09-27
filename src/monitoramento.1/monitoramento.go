package main

import "fmt"
import "os"
import "net/http"

func main() {

	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:

		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo os logs...")
		fmt.Println()
	case 3:
		fmt.Println("Saindo do programa...")
		os.Exit(0)

	default:
		fmt.Println("Escolha uma opção válida!")
		os.Exit(-1)
	}

}

func exibeMenu() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")

}

func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	// fmt.Println(resp)

	if resp.StatusCode == 200 {

	} else {

	}

}
