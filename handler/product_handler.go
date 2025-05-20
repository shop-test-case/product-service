package handler

import (
	"net/http"
	"product-service/controller"
	"product-service/model"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Ctrl *controller.ProductController
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.Ctrl.GetProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) AddProduct(c *gin.Context) {
	var p model.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.Ctrl.AddProduct(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "product created"})
}
