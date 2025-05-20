package repository

import (
	"product-service/model"

	"gorm.io/gorm"
)

type IProductRepo interface {
	FindAll() ([]model.Product, error)
	Create(product *model.Product) error
}

type ProductRepo struct {
	DB *gorm.DB
}

func (r *ProductRepo) FindAll() ([]model.Product, error) {
	var products []model.Product

	err := r.DB.Find(&products).Error

	return products, err
}

func (r *ProductRepo) Create(p *model.Product) error {
	return r.DB.Create(p).Error
}
