package main

import "fmt"

func main() {
	primeiroNumeroDigitado := lePrimeiroNumero()
	segundoNumeroDigitado := leSegundoNumero()
	total := primeiroNumeroDigitado + segundoNumeroDigitado

	fmt.Println("A soma dos dois números digitados é:", total)
	fmt.Println()

}

func lePrimeiroNumero() int {
	var numero int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&numero)
	return numero
}

func leSegundoNumero() int {
	var numero int
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&numero)
	return numero
}
