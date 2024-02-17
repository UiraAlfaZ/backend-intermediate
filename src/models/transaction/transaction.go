package transaction

import (
	"Backend/backend-api/src/config"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	IdProduct   string
	ProductName string
	Payment     int
}

func SelectAll() *gorm.DB {
	items := []Transaction{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB {
	var item Transaction
	return config.DB.First(&item, "id = ?", id)
}

func Post(bucket *Transaction) *gorm.DB {
	return config.DB.Create(&bucket)
}

func Updates(id string, newTransaction *Transaction) *gorm.DB {
	var updateTransaction Transaction
	return config.DB.Model(&updateTransaction).Where("id = ?", id).Updates(&newTransaction)
}

func Deletes(id string) *gorm.DB {
	var updateTransaction Transaction
	return config.DB.Delete(&updateTransaction, "id = ?", id)
}
