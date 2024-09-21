package main

import (
	_ "servergolang/docs"
	"servergolang/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/products", handlers.GetProducts)
	r.GET("/categories", handlers.GetCategories)

	r.Run(":8080")


}