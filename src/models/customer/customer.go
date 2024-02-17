package customer

import (
	"Backend/backend-api/src/config"

	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	IdCustomer string
	Name       string
	Address    string
}

func SelectAllOrder() *gorm.DB {
	newCustomer := []Customer{}
	return config.DB.Find(&newCustomer)
}

func Select(Id string) *gorm.DB {
	var profileCustomer Customer
	return config.DB.First(&profileCustomer, "id = ?", Id)
}

func Order(profileCustomer *Customer) *gorm.DB {
	return config.DB.Create(&profileCustomer)
}

func Updates(Id string, newCustomer *Customer) *gorm.DB {
	var updateCustomer Customer
	return config.DB.Model(&updateCustomer).Where("id = ?", Id).Updates(&newCustomer)
}

func Deletes(Id string) *gorm.DB {
	var updateCustomer Customer
	return config.DB.Delete(&updateCustomer, "id = ?", Id)
}
