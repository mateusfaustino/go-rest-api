package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusfaustino/go-rest-api/controller"
	"github.com/mateusfaustino/go-rest-api/db"
	usecase "github.com/mateusfaustino/go-rest-api/useCase"
)

func main() {
	server := gin.Default()

	_, err := db.ConnectDb()

	if err!= nil {
		panic(err)
	}

	productUseCase := usecase.NewProductUseCase()
	productController := controller.NewProductController(productUseCase)
	server.GET("/", func (ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})

	})

	server.GET("product", productController.GetProducts)

	server.Run(":8000")
}
