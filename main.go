package main

import (
	"log"
	"os"
	"projeto/appi"
)

func main() {

	aplicacao := appi.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
