package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fabiokusaba/devbook/api/src/config"
	"github.com/fabiokusaba/devbook/api/src/router"
)

func main() {
	config.Carregar()

	fmt.Println("Rodando API!")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
