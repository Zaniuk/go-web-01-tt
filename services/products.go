package services

import (
	"Zaniuk/clase-01-TT/repository"
	"Zaniuk/clase-01-TT/types"
)

type ProductServiceI interface {
	GetAll() []types.Product
	GetByID(id int) types.Product
	GetByParams(params map[string]string) []types.Product
	GetByPriceGreaterThan(price float64) []types.Product
}

type ProductService struct {
	repository repository.ProductRepository
}

func (p *ProductService) GetAll() []types.Product {
	return p.repository.GetAll()
}

func (p *ProductService) GetByID(id int) types.Product {
	return p.repository.GetByID(id)
}

func (p *ProductService) GetByParams(params map[string]string) []types.Product {
	return p.repository.GetByParams(params)
}

func (p *ProductService) GetByPriceGreaterThan(price float64) []types.Product {
	return p.repository.GetByPriceGreaterThan(price)
}
