package main

import "fmt"

func main() {

	fmt.Println()
	fmt.Println()
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
	fmt.Println()
	fmt.Println()

	var comando int

	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi:", comando)

	fmt.Println()

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo os logs...")
	// } else if comando == 3 {
	// 	fmt.Println("Saindo do programa...")
	// } else {
	// 	fmt.Println("Escolha uma opção válida!")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
		fmt.Println()
	case 2:
		fmt.Println("Exibindo os logs...")
		fmt.Println()
	case 3:
		fmt.Println("Saindo do programa...")
		fmt.Println()
	default:
		fmt.Println("Escolha uma opção válida!")
		fmt.Println()
	}

}
