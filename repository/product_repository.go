package repository

import (
	"database/sql"
	"log"

	"github.com/mateusfaustino/go-rest-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		log.Printf("Erro ao executar a query GetProducts: %v", err)
		return nil, err
	}
	defer rows.Close()

	var productList []model.Product

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			log.Printf("Erro ao escanear o produto: %v", err)
			return nil, err
		}
		productList = append(productList, product)
	}

	// Verifica por erros pós-iterações, como fechamento do cursor
	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração das rows: %v", err)
		return nil, err
	}

	return productList, nil
}
