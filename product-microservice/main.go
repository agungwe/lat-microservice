package main

import (
	"product/config"
	"product/product"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	r := gin.Default()
	db := config.DBInit()
	product := product.Product{DB: db}

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	config.RegisterConsul()
	config.RegisterZipkin()

	r.GET("/products", product.GetProducts)
	r.GET("/product/:id", product.GetProductById)
	r.POST("/products", product.CreateProduct)

	r.GET("/healthcheck", config.Healthcheck)

	r.Run() //port 8080
}
