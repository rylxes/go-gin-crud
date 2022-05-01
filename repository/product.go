package repository

import (
	"goGin/config"
	"goGin/entities"
	"gorm.io/gorm"
	"log"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepostiory(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() []entities.Product {
	var products []entities.Product

	db := config.InitDB()
	log.Println("db -> {}", db.Name)
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) FindByID(id uint) entities.Product {
	var product entities.Product
	p.DB.First(&product, id)

	return product
}

func (p *ProductRepository) Save(product entities.Product) entities.Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) Delete(product entities.Product) {
	p.DB.Delete(&product)
}
