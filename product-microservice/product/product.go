package product

import (
	"math/rand"
	"product/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.opencensus.io/trace"
)

type Product struct {
	DB *gorm.DB
}

type product struct {
	ProductName string `json:"product_name"`
	SKU         string `json:"sku"`
	Qty         int32  `json:"qty"`
}

func (p *Product) GetProducts(c *gin.Context) {
	db := p.DB
	var products []model.Product

	db.Find(&products)

	c.JSON(200, gin.H{
		"data": products,
	})

	serviceb(c)
}

func (p *Product) GetProductById(c *gin.Context) {
	var product model.Product
	db := p.DB
	id := c.Param("id")

	db.Where("id = ?", id).Find(&product)

	c.JSON(200, gin.H{
		"product_name": product.ProductName,
		"sku":          product.SKU,
		"qty":          product.Qty,
	})
}

func (p *Product) CreateProduct(c *gin.Context) {
	var request product

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	p.DB.Create(&request)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func serviceb(c *gin.Context) {
	_, span := trace.StartSpan(c, "/products")
	defer span.End()
	time.Sleep(time.Duration(rand.Intn(800)+200) * time.Millisecond)
}
