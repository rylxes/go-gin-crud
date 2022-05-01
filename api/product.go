package api

import (
	"github.com/gin-gonic/gin"
	"goGin/dto"
	"goGin/entities"
	"goGin/mappers"
	"goGin/service"
	"log"
	"net/http"
	"strconv"
)

type ProductAPI struct {
	ProductService service.ProductService
}

func ProvideProductAPI(p service.ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

func (p *ProductAPI) FindAll(c *gin.Context) {
	log.Println("Entered!!")
	products := p.ProductService.FindAll()

	c.JSON(http.StatusOK, gin.H{"products": mappers.ToProductDTOs(products)})
}

func (p *ProductAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": mappers.ToProductDTO(product)})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var productDTO dto.ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdProduct := p.ProductService.Save(mappers.ToProduct(productDTO))

	c.JSON(http.StatusOK, gin.H{"product": mappers.ToProductDTO(createdProduct)})
}

func (p *ProductAPI) Update(c *gin.Context) {
	var productDTO dto.ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product == (entities.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	product.Code = productDTO.Code
	product.Price = productDTO.Price
	p.ProductService.Save(product)

	c.Status(http.StatusOK)
}

func (p *ProductAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product == (entities.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.Delete(product)

	c.Status(http.StatusOK)
}
