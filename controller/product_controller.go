package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusfaustino/go-rest-api/model"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context){
	products:= []model.Product{
		{
			Id: 1,
			Name: "relógio",
			Price: 100,
		},
		{
			Id: 2,
			Name: "óculos",
			Price: 73.00,
		},
		{
			Id: 1,
			Name: "boné",
			Price: 55.44,
		},
	}

	ctx.JSON(http.StatusOK, products)
}