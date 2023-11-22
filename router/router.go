package router

import (
	"os"
	"products/controller/product_controller"
	"products/db"

	"github.com/gin-gonic/gin"
)

// var PORT = ":8080"

func init() {
	db.InitializeDB()
}

func StartRouter() {
	router := gin.Default()

	var PORT = os.Getenv("PORT")
	
	productRouter := router.Group("/products")
	{
		productRouter.POST("/", product_controller.CreateProduct)
		productRouter.PUT("/:productId", product_controller.UpdateProduct)
		productRouter.GET("/", product_controller.GetProducts)
		productRouter.DELETE("/:productId", product_controller.DeleteProduct)
	}

	router.Run(":" +PORT)
}
