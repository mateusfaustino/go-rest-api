package usecase

import "github.com/mateusfaustino/go-rest-api/model"

type ProductUseCase struct {
	//Repository
}

func NewProductUseCase() ProductUseCase {
	return ProductUseCase{}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product,error){
	return []model.Product{}, nil
}