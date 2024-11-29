package router
import (
	gin "github.com/gin-gonic/gin"
)
func Initialize(){
	//inicializa o router utilizando as configurações Default do gin
	router := gin.Default()
//cria uma rota para o endpoint /ping
	router.GET("/ping", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
//rodando a API
	router.Run(":8080")
}