package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusfaustino/go-rest-api/controller"
)

func main() {
	server := gin.Default()


	productController := controller.NewProductController()
	server.GET("/", func (ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})

	})

	server.GET("product", productController.GetProducts)

	server.Run(":8000")
}
