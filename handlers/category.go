package handlers

import (
	"net/http"
	"servergolang/database"
	"servergolang/models"

	"github.com/gin-gonic/gin"
)

// GetCategories godoc
// @Summary Get all categories
// @Description Get a list of all categories
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	var categories []models.Category
	database.DB.Preload("Products").Find( &categories )
	c.JSON(http.StatusOK, categories)
}


// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category in the database
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category to create"
// @Success 201 {object} models.Category
// @Failure 400 {object} gin.H{"error": string}
// @Router /categories [post]
func CreateCategory(	c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&category)
	c.JSON( http.StatusOK, category )
}