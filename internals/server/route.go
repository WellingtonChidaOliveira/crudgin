package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wellingtonchida/products-with-gin/types"
)

func (s *server) NewRouter() http.Handler {
	r := gin.Default()

	r.GET("/ping", s.HandlerPing)
	r.GET("/products", s.HandleGetProducts)
	r.GET("/health", s.HandlerHealthCheck)
	r.GET("/products/:id", s.HandleGetProductByID)
	r.POST("/products", s.HandleCreateProduct)

	return r
}

func (s *server) HandlerPing(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *server) HandlerHealthCheck(ctx *gin.Context) {
	msgs := s.db.HealthCheck()
	ctx.JSON(200, gin.H{
		"healthy": msgs,
	})
}

func (s *server) HandleGetProducts(ctx *gin.Context) {
	products, err := s.db.GetProducts()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"products": products,
	})
}

func (s *server) HandleGetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := s.db.GetProductByID(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"product": product,
	})
}

func (s *server) HandleCreateProduct(ctx *gin.Context) {
	var product types.ProductRequest
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad Request",
		})
	}
	if err := s.db.CreateProduct(product); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(201, gin.H{
		"message": "Product created",
	})
}

func (s *server) HandleUpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product types.ProductRequest
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad Request",
		})
	}
	if err := s.db.UpdateProduct(id, product); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Product updated",
	})
}

func (s *server) HandleDeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := s.db.DeleteProduct(id); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Product deleted",
	})
}
