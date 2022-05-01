package service

import (
	"goGin/entities"
	"goGin/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func ProvideProductService(p repository.ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) FindAll() []entities.Product {
	return p.ProductRepository.FindAll()
}

func (p *ProductService) FindByID(id uint) entities.Product {
	return p.ProductRepository.FindByID(id)
}

func (p *ProductService) Save(product entities.Product) entities.Product {
	p.ProductRepository.Save(product)

	return product
}

func (p *ProductService) Delete(product entities.Product) {
	p.ProductRepository.Delete(product)
}
