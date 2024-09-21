package handlers

import (
	"net/http"
	"servergolang/models"

	"github.com/gin-gonic/gin"
)

var products = []models.Product{
	{ID: 1, Name: "Product 1", Description: "This is product 1", Price: 9.99, Quantity: 10, CategoryID: 1},
	{ID: 2, Name: "Product 2", Description: "This is product 2", Price: 19.99, Quantity: 20, CategoryID: 2},
	{ID: 3, Name: "Product 3", Description: "This is product 3", Price: 29.99, Quantity: 30, CategoryID: 3},
}


// GetProducts godoc
// @Summary Get all products
// @Description Get a list of all products
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products);
}