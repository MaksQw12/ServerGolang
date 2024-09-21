package handlers

import (
	"net/http"
	"servergolang/models"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{
	{ID: 1, Name: "Category 1", },
	{ID: 2, Name: "Category 2", },
	{ID: 3, Name: "Category 3", },	
}


// GetCategories godoc
// @Summary Get all categories
// @Description Get a list of all categories
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories);
}