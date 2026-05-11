package useCase

import (
	"api-product/model"
	"database/sql"
)

type ProductUseCase struct {
	DB *sql.DB
}

func NewProductUseCase(db *sql.DB) ProductUseCase {
	return ProductUseCase{DB: db}
}

func (uc *ProductUseCase) GetProducts() ([]model.Product, error) {
	rows, err := uc.DB.Query("SELECT id, name, price FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
