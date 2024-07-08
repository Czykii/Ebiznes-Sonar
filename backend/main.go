package main

import (
	"backend/controllers"
	"backend/database"
	"backend/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World!")
	})

	DB := database.Connect()

	productData := []models.Product{
		{Name: "Produkt 1", Description: "Lorem ipsum", Price: 350.99},
		{Name: "Produkt 2", Description: "Lorem ipsum", Price: 55.99},
		{Name: "Produkt 3", Description: "Lorem ipsum", Price: 1250.0},
	}
	
	err := database.GetData(DB, productData)
	if err != nil {
		panic(err)
	}

	productController := controllers.CreateProductController(DB)
	e.GET("/products", productController.GetProducts)
	paymentController := controllers.NewPaymentController(DB)
	e.POST("/payment", paymentController.MakePayment)

	e.Logger.Fatal(e.Start(":8080"))
}
