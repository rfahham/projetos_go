package main

import (
	"fmt"
)

func main() {

	var precoLeite, precoOvo, precoPao, valorPago float32

	fmt.Println("Informe o preço do leite:")
	fmt.Scanln(&precoLeite)
	fmt.Println("Informe o preço do ovo:")
	fmt.Scanln(&precoOvo)
	fmt.Println("Informe o preço do Pão:")
	fmt.Scanln(&precoPao)

	var precoTotal = precoLeite + precoOvo + precoPao

	fmt.Println()
	fmt.Println("O preço do leite é R$", precoLeite, ",")
	fmt.Println("O preço do ovo é R$", precoOvo, ",")
	fmt.Println("O preço do pão é R$", precoPao, ",")
	fmt.Println()
	fmt.Println("O preço total das suas compras é: ", precoTotal, ",")
	fmt.Println()

	fmt.Println("Informe o valor pago:")
	fmt.Scanln(&valorPago)

	var valorTroco = precoTotal - valorPago

	fmt.Println()
	fmt.Print("Seu troco R$ ", valorTroco)
	fmt.Println()

}
