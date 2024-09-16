package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Tasks struct{
	Id int `json:"id"`
	Title string `json:"title"`
}


//slice
var taskList = []Tasks{
	{Id:1, Title:"Estudar"},
	{Id:2, Title: "Caminhar"},
}


func RouteTest(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"primeira api"})
}

func ListarTarefas(ctx *gin.Context){
	ctx.IndentedJSON(http.StatusOK, taskList)
}

func ListarTarefasPorId(ctx *gin.Context){
	//armazenando o valor do parametro da url com ctx.Param
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam) //conversao do valor de parametro 
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"id invalido"})
	}
	for _, task := range taskList{
		if (task.Id == id){
				ctx.JSON(http.StatusOK, task)
				return
		}
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"error":"Tarefa não encontrada"})
}

func AdicionarTarefas(ctx *gin.Context){
	var newTask Tasks
		if err := ctx.BindJSON(&newTask); err !=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		newTask.Id = len(taskList) + 1
		taskList = append(taskList, newTask)
		ctx.JSON(http.StatusCreated, gin.H{"message":"Item criado com sucesso","item":taskList})
}

func DeletarTarefa(ctx *gin.Context){
	idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam) //conversao do valor de parametro 
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"Error":"id inválido"})
		}
		for i, task := range taskList{
			if (task.Id == id){
				taskList = append(taskList[:i], taskList[i+1:]...)
				ctx.JSON(200, gin.H{"Message":"Deletado com sucesso"})
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message":"Tarefa não encontrada"})
}

func AlterarTarefa(ctx *gin.Context){
	idParam := ctx.Param("id")
		id,err := strconv.Atoi(idParam)
		var upadatedTask Tasks
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error":"id invalido"})
		}
		if err := ctx.BindJSON(&upadatedTask); err !=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		for index, task := range taskList{
			upadatedTask.Id = task.Id
			if (task.Id == id){
				taskList[index] = upadatedTask //trocando as informações consultadas por index pelo updatedTask
				ctx.JSON(http.StatusOK, gin.H{"Message":"Item atualizado", "item":upadatedTask})
			}
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message":"Tarefa não encontrada"})
}