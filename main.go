package main

import (
	"github.com/WillJR183/ci-cd-example/database"
	"github.com/WillJR183/ci-cd-example/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
