package main

import (
	"net/http"

	"github.com/ngdlong91/funtech/cmd/gin/pkg/product"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"state": "running",
		})
	})

	productService := product.NewGinService()
	r.POST("purchase", productService.Purchase)

	if err := r.Run(":3000"); err != nil {
		panic("cannot running server ")
	}
}
