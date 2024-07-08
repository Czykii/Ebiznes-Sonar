package database

import (
	"backend/models"
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("data.db"))
	if err != nil {
		panic("Failed to connect to database")
	}

	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		panic("Failed to migrate database")
	}

	return DB
}

func GetData(db *gorm.DB, data []models.Product) error {
	for _, product := range data {
		var existingProduct models.Product
		res := db.Where("name = ? AND description = ?", product.Name, product.Description).First(&existingProduct)
		if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			if err := db.Create(&product).Error; err != nil {
				return err
			}
		}
	}

	return nil
}