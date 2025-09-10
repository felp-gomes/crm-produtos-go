package main

import (
	"fmt"
	"net/http"
	"produtos/routes"

	_ "github.com/lib/pq"
)


func main() {
	fmt.Println("Servidor rodando inicializado: http://localhost:8000")
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
