package seller

import (
	"Backend/backend-api/src/config"

	"github.com/jinzhu/gorm"
)

type Seller struct {
	gorm.Model
	IdSeller string
	Name     string
}

func SelectAll() *gorm.DB {
	items := []Seller{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB {
	var item Seller
	return config.DB.First(&item, "id = ?", id)
}

func Post(profileSeller *Seller) *gorm.DB {
	return config.DB.Create(&profileSeller)
}

func Updates(id string, newSeller *Seller) *gorm.DB {
	var updateSeller Seller
	return config.DB.Model(&updateSeller).Where("id = ?", id).Updates(&newSeller)
}

func Deletes(id string) *gorm.DB {
	var updateSeller Seller
	return config.DB.Delete(&updateSeller, "id = ?", id)
}
