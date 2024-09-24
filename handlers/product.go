package handlers

import (
	"net/http"
	"servergolang/database"
	"servergolang/models"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary Get all products
// @Description Get a list of all products
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Preload("Category").Find(&products)
	c.JSON(http.StatusOK, products)
}


// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product in the database
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product to create"
// @Success 201 {object} models.Product
// @Failure 400 {object} gin.H{"error": string}
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}