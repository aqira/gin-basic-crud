package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductsRegister(router *gin.RouterGroup) {
	router.GET("/:id", retrieveProduct)
	router.GET("/", getProductList)
	router.POST("/", createProduct)
	router.PUT("/:id", updateProduct)
	router.DELETE("/:id", ProductDelete)
}

func createProduct(c *gin.Context) {
	var newProduct Product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	newProduct.Id = currentId
	productsMap[currentId] = newProduct
	currentId++

	c.IndentedJSON(http.StatusCreated, newProduct)
}

func getProductList(c *gin.Context) {

	var productsList []Product
	for _, product := range productsMap {
		productsList = append(productsList, product)
	}

	c.IndentedJSON(http.StatusOK, productsList)
}

func retrieveProduct(c *gin.Context) {
	convertedId, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(
			http.StatusForbidden,
			gin.H{"message": "invalid Id input, must be a number"},
		)
		return
	}

	id := uint(convertedId)
	product := productsMap[id]

	c.IndentedJSON(http.StatusOK, product)
}

func updateProduct(c *gin.Context) {
		convertedId, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(
			http.StatusForbidden,
			gin.H{"message": "invalid Id input, must be a number"},
		)
		return
	}

	id := uint(convertedId)

	var updatedProduct Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	if _, exists := productsMap[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	updatedProduct.Id = id
	productsMap[id] = updatedProduct
	c.IndentedJSON(http.StatusOK, updatedProduct)
}

func ProductDelete(c *gin.Context) {
	convertedId, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(
			http.StatusForbidden,
			gin.H{"message": "invalid Id input, must be a number"},
		)
		return
	}

	id := uint(convertedId)
	delete(productsMap, id)

	c.Status(http.StatusNoContent)
}
