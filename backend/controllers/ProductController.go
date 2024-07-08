package controllers

import (
	"backend/models"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type ProductController struct {
	DB *gorm.DB
}

func CreateProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

func (productControler *ProductController) GetProducts(context echo.Context) error {
	var products []models.Product
	if err := productControler.DB.Find(&products).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return context.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return context.JSON(http.StatusOK, products)
}