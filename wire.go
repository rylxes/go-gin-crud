package main

import (
	"goGin/api"
	"goGin/repository"
	"goGin/service"
	"gorm.io/gorm"
)

func InitProductAPI(db *gorm.DB) api.ProductAPI {
	productRepository := repository.ProvideProductRepostiory(db)
	productService := service.ProvideProductService(productRepository)
	productAPI := api.ProvideProductAPI(productService)
	return productAPI
}
