package main

import (
	"ecommerce/database"
	"ecommerce/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()

}
