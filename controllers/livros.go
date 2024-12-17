package controllers

import (
	"Projeto-da-Bibilioteca/database"
	"Projeto-da-Bibilioteca/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeLivros(c *gin.Context) {
	var livros []models.Livro
	database.DB.Find(&livros)
	c.JSON(200, livros)
}

func BuscaLivroPorId(c *gin.Context) {
	var livro models.Livro
	id := c.Params.ByName("id")
	database.DB.First(&livro, id)

	if livro.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, livro)
}

func BuscarNomeDoLivro(c *gin.Context, nome string) {

}
