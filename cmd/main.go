package main

import (
	"log"

	"github.com/mateusfaustino/go-rest-api/controller"
	"github.com/mateusfaustino/go-rest-api/db"
	"github.com/mateusfaustino/go-rest-api/repository"
	"github.com/mateusfaustino/go-rest-api/routes"
	"github.com/mateusfaustino/go-rest-api/usecase"
)

func main() {

	connection, err := db.ConnectDb()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}

	ProductRepository := repository.NewProductRepository(connection)
	productUseCase := usecase.NewProductUseCase(ProductRepository)
	productController := controller.NewProductController(productUseCase)

	// Configurar rotas e iniciar o servidor
	server := routes.SetupRouter(productController)
	server.Run(":8000")
}
