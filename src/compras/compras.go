package main

import "fmt"

func main() {
	var precoLeite float32 = 2.99
	var precoOvo float32 = 3.99
	var precoPao float32 = 0.99

	var precoTotal = precoLeite + precoOvo + precoPao
	fmt.Println()
	fmt.Println("O preço do leite é R$", precoLeite)
	fmt.Println("O preço do ovo é R$", precoOvo)
	fmt.Println("O preço do pão é R$", precoPao)
	fmt.Println()
	fmt.Println("O total da compra é R$", precoTotal)
	fmt.Println()
}
