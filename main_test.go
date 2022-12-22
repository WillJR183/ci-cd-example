package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/WillJR183/ci-cd-example/controllers"
	"github.com/WillJR183/ci-cd-example/database"
	"github.com/WillJR183/ci-cd-example/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Auno Teste",
		RG:   "123456789",
		CPF:  "12345678912",
	}

	database.DB.Create(&aluno)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDeSaudacoesComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacoes)
	request, _ := http.NewRequest("GET", "/will", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")
	responseMock := `{"API diz": "E ae will, tudo jóia?"}`
	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, responseMock, string(responseBody))
}

func TestListaTodosAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.TodosAlunos)
	request, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestBuscaAlunoPorCpfHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoPorCpf)
	request, _ := http.NewRequest("GET", "/alunos/cpf/12345678912", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestBuscaAlunoPorIdHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscarAlunoPorId)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("GET", pathDaBusca, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var alunoMock models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Auno Teste", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, "12345678912", alunoMock.CPF)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("DELETE", pathDaBusca, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditarAluno)

	aluno := models.Aluno{
		Nome: "Aluno Teste",
		RG:   "123567891",
		CPF:  "13579246802",
	}

	valorJson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, "Aluno Teste", &alunoMockAtualizado.Nome)
	assert.Equal(t, "123567891", &alunoMockAtualizado.RG)
	assert.Equal(t, "13579246802", &alunoMockAtualizado.CPF)
}
