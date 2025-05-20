package main

import (
	"product-service/config"
	"product-service/controller"
	"product-service/database"
	"product-service/handler"
	"product-service/middleware"
	"product-service/model"
	"product-service/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)
	db.AutoMigrate(&model.Product{})

	repo := &repository.ProductRepo{DB: db}
	ctrl := &controller.ProductController{Repo: repo}
	h := &handler.ProductHandler{Ctrl: ctrl}

	r := gin.Default()
	r.Use(cors.Default())

	auth := r.Group("/")
	auth.Use(middleware.JWT(cfg.JWTSecret))
	auth.GET("/products", h.GetProducts)
	auth.POST("/products", h.AddProduct)

	r.Run(":" + cfg.Port)
}
