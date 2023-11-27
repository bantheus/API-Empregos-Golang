package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//* Inicializa o router utilizando as configurações Default do gin
	r := gin.Default()

	//* Define a rota /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//* Servidor rodando em localhost:8080
	r.Run(":8080")
}
