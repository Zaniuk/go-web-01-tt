package main

// import gin

import (
	"Zaniuk/clase-01-TT/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Greeting struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func main() {
	// ...

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/products", func(c *gin.Context) {
		var productService services.ProductServiceI = &services.ProductService{}
		products := productService.GetAll()

		c.JSON(200, products)

	})

	router.GET("/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid id"})
			return
		}
		productService := services.ProductService{}
		product := productService.GetByID(id)
		if product.ID == 0 {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(200, product)
	})

	router.GET("/products/search", func(c *gin.Context) {
		var productService services.ProductServiceI = &services.ProductService{}
		param := c.Query("price")

		price, err := strconv.ParseFloat(param, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid price"})
			return
		}

		products := productService.GetByPriceGreaterThan(price)

		c.JSON(200, products)
	})
	router.Run(":80") // listen and serve on
}
