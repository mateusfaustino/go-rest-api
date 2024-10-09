package usecase

import (
	"errors"

	"github.com/mateusfaustino/go-rest-api/model"
	"github.com/mateusfaustino/go-rest-api/repository"
)

type ProductUseCase struct {
	Repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		Repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	products, err := pu.Repository.GetProducts()
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("nenhum produto encontrado")
	}

	return products, nil
}
