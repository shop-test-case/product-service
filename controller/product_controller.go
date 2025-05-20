package controller

import (
	"product-service/model"
	"product-service/repository"
)

type ProductController struct {
	Repo repository.IProductRepo
}

func (c *ProductController) GetProducts() ([]model.Product, error) {
	return c.Repo.FindAll()
}

func (c *ProductController) AddProduct(p *model.Product) error {
	return c.Repo.Create(p)
}
