package main

import (
	"github.com/gin-gonic/gin"
)



func main(){
	router := gin.Default()
	RegisterRoutes(router)
	
	//recusa proxys
	router.SetTrustedProxies(nil)
	router.Run(":8000")
}

