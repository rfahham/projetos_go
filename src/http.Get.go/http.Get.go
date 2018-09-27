package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://www.globo.com/economia/noticia/nao-ha-possibilidade-de-criacao-de-imposto-para-financiar-seguranca-diz-meirelles.ghtml")
	if err != nil {
		log.Fatal(err)
	}

	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")
	} else {
		fmt.Println("Argh! Broken")
	}
}
