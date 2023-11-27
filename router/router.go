package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//* Inicializa o router
	r := gin.Default()

	//* Inicializa as rotas
	initializeRoutes(r)

	//* Servidor rodando em localhost:8080
	r.Run(":8080")
}
