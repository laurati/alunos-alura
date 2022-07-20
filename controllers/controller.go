package controllers

import (
	"ecommerce/database"
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ExibeAluno
// @Summary Exibe Alunos
// @Description Exibe Alunos
// @Tags Alunos
// @Accept json
// @Produce json
// @Success 200 {array} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	// Failure 400  {object} httputil.HTTPError
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

// CriaAluno
// @Summary Cria Alunos
// @Description Cria Alunos
// @Tags Alunos
// @Accept json
// @Produce json
// @Param        data  body   models.Aluno   true  "Body para criar um aluno"
// @Success 200 {object} models.Aluno
// @Failure 400  {object} httputil.HTTPError
// @Failure 404 "Not Found"
// @Failure 500 "Internal Server Error"
// @Router /alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// ExibeAluno
// @Summary Exibe Alunos por ID
// @Description Exibe Alunos por ID
// @Tags Alunos
// @Accept json
// @Produce json
// @Param        id   path      int  true  "Aluno ID"
// @Success 200 {object} models.Aluno
// @Failure 400 "Bad Request"
// @Failure 404 "Not Found"
// @Failure 500 "Internal Server Error"
// @Router /alunos/{id} [get]
func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno
// @Summary Deleta Alunos
// @Description Deleta Alunos
// @Tags Alunos
// @Accept json
// @Produce json
// @Param id path int true "Aluno ID"
// @Success 200 {object} models.Aluno
// @Failure 400  {object} httputil.HTTPError
// @Failure 404 "Not Found"
// @Failure 500 "Internal Server Error"
// @Router /alunos/{id} [delete]
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

// EditaAluno
// @Summary Edita Alunos
// @Description Edita Alunos
// @Tags Alunos
// @Accept json
// @Produce json
// @Param        id   path      int  true  "Aluno ID"
// @Param        data  body   models.Aluno   true  "O body para editar um aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 "Bad Request"
// @Failure 404 "Not Found"
// @Failure 500 "Internal Server Error"
// @Router /alunos/{id} [patch]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Save(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// ExibeAluno
// @Summary Exibe Alunos por CPF
// @Description Exibe Alunos por CPF
// @Tags Alunos
// @Accept json
// @Produce json
// @Param        cpf   path      string  true  "CPF do aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 "Bad Request"
// @Failure 404 "Not Found"
// @Failure 500 "Internal Server Error"
// @Router /alunos/cpf/{cpf} [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
