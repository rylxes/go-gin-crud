package mappers

import (
	"goGin/dto"
	"goGin/entities"
)

func ToProduct(productDTO dto.ProductDTO) entities.Product {
	return entities.Product{Code: productDTO.Code, Price: productDTO.Price}
}

func ToProductDTO(product entities.Product) dto.ProductDTO {
	return dto.ProductDTO{ID: product.ID, Code: product.Code, Price: product.Price}
}

func ToProductDTOs(products []entities.Product) []dto.ProductDTO {
	productdtos := make([]dto.ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}
