package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const monitoramentos = 1

func main() {

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarVerificacao()
		case 2:
			fmt.Println("Lista de URLS Status Code 200...")
			fmt.Println("")
			imprimeLista200()
		case 3:
			fmt.Println("Lista de URLS Status Code 404...")
			fmt.Println("")
			imprimeLista404()
		case 4:
			fmt.Println("Saindo do programa... Obrigado!")
			fmt.Println("")
			os.Exit(0)
		default:
			fmt.Println("Selecione uma opção válida!")
			fmt.Println("")
			os.Exit(-1)
		}
	}

}

func exibeMenu() {
	fmt.Println("")
	fmt.Println("Programa para separar URLS que não sejam Status Code 200.")
	fmt.Println("")
	fmt.Println("Escolha uma das opções abaixo:")
	fmt.Println("")
	fmt.Println("1 - Executar Programa")
	fmt.Println("2 - Lista de aquivos Status Code 200")
	fmt.Println("3 - Lista de aquivos Status Code 404")
	fmt.Println("4 - Sair do Programa")
	fmt.Println("")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("A opção escolhida foi:", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarVerificacao() {
	fmt.Println("Verificando URLS...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println(i)
			testaSite(site)
		}

		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, "Status Code: 200")
		registraLog200(site, true)
	} else {
		fmt.Println(site, "Status Code:", resp.StatusCode)
		registraLog404(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("lista.csv")

	if err != nil {
		fmt.Println("Status do arquivo:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}

func registraLog200(site string, status bool) {

	arquivo, err := os.OpenFile("200.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(site + "\n")

	arquivo.Close()
}

func registraLog404(site string, status bool) {

	arquivo, err := os.OpenFile("404.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(site + "\n")

	arquivo.Close()
}

func imprimeLista200() {

	arquivo, err := ioutil.ReadFile("200.csv")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}

func imprimeLista404() {

	arquivo, err := ioutil.ReadFile("404.csv")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
