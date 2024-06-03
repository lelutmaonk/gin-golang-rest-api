package main

import (
	productcontroller "go_tutorial/rest_api_gin/controller/productController"
	"go_tutorial/rest_api_gin/model"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	model.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:product_id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:product_id", productcontroller.Update)
	r.DELETE("/api/products/:product_id", productcontroller.Delete)

	r.Run()

}
