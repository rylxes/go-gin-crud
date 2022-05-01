package main

import (
	"github.com/gin-gonic/gin"
	"goGin/config"
	"log"
)

const port = 8000

func main() {
	db := config.InitDB()
	log.Println("db -> {}", db.Name)
	productAPI := InitProductAPI(db)

	r := gin.Default()

	log.Println("db -> {}", db.Name)
	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
