package repository

import (
	"database/sql"

	_"github.com/mateusfaustino/go-rest-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection * sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

// func (pr *ProductRepository) GetProducts() ([]model.Product, error){
// 	query := "SELECT id, name, price FROM products"

// }