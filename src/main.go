package main

import (
	"match-dos-estudos/src/database"
	"match-dos-estudos/src/router"
)

// @title           Match dos Estudos API
// @version         1.0
// @description     API REST para match de estudantes com sessões de estudo
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// 1. Conecta ao MongoDB
	database.Connect()
	defer database.Disconnect()

	// 2. Insere dados de exemplo (só executa se as coleções estiverem vazias)
	database.Seed()

	// 3. Sobe o servidor HTTP na porta 8080
	r := router.SetupRouter()
	r.Run(":8080")
}
