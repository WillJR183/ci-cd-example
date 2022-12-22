package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/WillJR183/ci-cd-example/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/")
	r.Static("/assets", "./assets")
	r.GET("/:nome", controllers.Saudacoes)
	r.GET("/alunos", controllers.TodosAlunos)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoPorCpf)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}
