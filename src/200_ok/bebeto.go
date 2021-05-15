package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		urlHost string
		origem  string
		destino string
	)

	flag.StringVar(&urlHost, "url", "", "url completa. exemplo: http://globoesporte.qa02.globoi.com/")
	flag.StringVar(&origem, "orig", "./origem.txt", "caminho do arquivo de origem")
	flag.StringVar(&destino, "dst", "./destino.txt", "caminho do arquivo de destino")

	flag.Parse()

	if urlHost == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fileRead, err := os.Open(origem)
	if err != nil {
		log.Fatal(err)
	}
	defer fileRead.Close()

	fileWrite, err := os.Create(destino)
	if err != nil {
		log.Fatal(err)
	}
	defer fileWrite.Close()

	scanner := bufio.NewScanner(fileRead)
	writer := bufio.NewWriter(fileWrite)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	count := []int{0, 0}

	for scanner.Scan() {

		count[0]++
		resposta, err := client.Get(urlHost + scanner.Text())
		if err != nil {
			log.Println(err.Error())
			continue
		}
		defer resposta.Body.Close()

		if resposta.StatusCode == 200 {
			count[1]++
			_, err := writer.WriteString(scanner.Text() + "\n")
			if err != nil {
				log.Println(err)
			}
		}
	}
	writer.Flush()

	readStat, _ := fileRead.Stat()
	writeStat, _ := fileWrite.Stat()

	fmt.Printf("Done\nRead: %d %d\nWrite: %d %d\n", readStat.Size(), count[0], writeStat.Size(), count[1])
}
