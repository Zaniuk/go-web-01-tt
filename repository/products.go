package repository

import (
	"Zaniuk/clase-01-TT/types"
)

type ProductRepositoryI interface {
	GetAll() []types.Product
	GetByID(id int) types.Product
	GetByParams(params map[string]string) []types.Product
	GetByPriceGreaterThan(price float64) []types.Product
}

type ProductRepository struct{}

func (p *ProductRepository) GetAll() []types.Product {
	return products
}

func (p *ProductRepository) GetByID(id int) types.Product {
	for _, product := range products {
		if product.ID == id {
			return product
		}
	}

	return types.Product{}
}

func (p *ProductRepository) GetByParams(params map[string]string) []types.Product {
	var result []types.Product

	for _, product := range products {
		if product.Name == params["name"] {
			result = append(result, product)
		}
	}

	return result
}

func (p *ProductRepository) GetByPriceGreaterThan(price float64) []types.Product {
	var result []types.Product

	for _, product := range products {
		if product.Price > price {
			result = append(result, product)
		}
	}

	return result
}

var products = []types.Product{
	{
		ID:          1,
		Name:        "Beef - Ox Tongue, Pickled",
		Quantity:    63,
		CodeValue:   "33261-011",
		IsPublished: true,
		Expiration:  "12-01-2023",
		Price:       213,
	},
	{
		ID:          2,
		Name:        "Wine - Alicanca Vinho Verde",
		Quantity:    3,
		CodeValue:   "59535-8101",
		IsPublished: true,
		Expiration:  "08-02-2023",
		Price:       402,
	},
	{
		ID:          3,
		Name:        "Bread - Ciabatta Buns",
		Quantity:    80,
		CodeValue:   "36987-1087",
		IsPublished: true,
		Expiration:  "13-06-2022",
		Price:       238,
	},
}
