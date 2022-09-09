package main

import (
	"fmt"

	"github.com/MarcosRoch4/api-rest-go/database"
	"github.com/MarcosRoch4/api-rest-go/routes"
)

func main() {

	database.ConectaDB()
	fmt.Println("Iniciando o servidor Rest com GO!")
	routes.HandleRequest()
}
