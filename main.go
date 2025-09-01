package main

import (
	"fmt"
	"net/http"
	"produtos/routes"

	_ "github.com/lib/pq"
)


func main() {
	fmt.Println("bola")
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
