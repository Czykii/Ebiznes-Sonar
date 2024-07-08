package controllers

import (
	"net/http"
	"strconv"
	"time"
	"backend/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PaymentController struct {
	DB *gorm.DB
}

func NewPaymentController(db *gorm.DB) *PaymentController {
	return &PaymentController{DB: db}
}

func (paymentcontroller *PaymentController) MakePayment(context echo.Context) error {
	var payment models.Payment
	if err := context.Bind(&payment); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	amount, err := strconv.ParseFloat(payment.Amount.String(), 64)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payment amount"})
	}

	if amount <= 0 {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payment amount"})
	}

	if payment.CardNumber == "" || payment.ExpiryDate == "" || isExpired(payment.ExpiryDate) {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Bad card"})
	}
	return context.JSON(http.StatusOK, map[string]string{"message": "Payment successful"})
}

func isExpired(dateStr string) bool {
	parsedDate, err := time.Parse("01/06", dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return true
	}
	currentDate := time.Now()
	return parsedDate.Before(currentDate)
}