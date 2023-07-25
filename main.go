package main

import (
	"tutorial/gin/crud/products"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api")
	products.ProductsRegister(v1.Group("/products"))

	r.Run()
}
