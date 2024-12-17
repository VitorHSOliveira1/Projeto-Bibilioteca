package routes

import (
	"Projeto-da-Bibilioteca/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/")
	r.GET("/products", controllers.ExibeLivros)
	r.GET("/:id", controllers.BuscaLivroPorId)
	r.Run(":5500")
}
