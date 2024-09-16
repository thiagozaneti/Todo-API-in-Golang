package main

import (
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(router *gin.Engine){
	router.GET("/", RouteTest)
	router.GET("/listarTarefas", ListarTarefas)
	router.GET("/listarTarefas/:id", ListarTarefasPorId)
	router.POST("/adicionarTarefas", AdicionarTarefas)
	router.POST("/deletarTarefa/:id", DeletarTarefa)
	router.PUT("/alterarTarefa/:id", AlterarTarefa)
}