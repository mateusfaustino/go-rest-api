package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusfaustino/go-rest-api/controller"
)

// SetupRouter configura as rotas da API e retorna o router configurado.
func SetupRouter(productController controller.ProductController) *gin.Engine {
	router := gin.Default()

	// Rota de exemplo para testar o servidor
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Grupo de rotas para produtos, com prefixo `/api`
	api := router.Group("/api")
	{
		api.GET("/product", productController.GetProducts)
	}

	return router
}
