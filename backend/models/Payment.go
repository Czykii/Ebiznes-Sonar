package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CardNumber string      `json:"cardNumber"`
	ExpiryDate string      `json:"expiryDate"`
	Amount     json.Number `json:"amount"`
}