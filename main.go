package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Configurando um manipulador (handler) para a rota "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Bem-vindo ao backend simples em Go!")
	})

	// Iniciando o servidor na porta 8080
	port := 8080
	fmt.Printf("Servidor iniciado em http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
