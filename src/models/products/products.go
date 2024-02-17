package products

import (
	"Backend/backend-api/src/config"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	IdItem         string
	ProductName    string
	DateProduction string
	Condition      string
	Stock          int
}

func SelectAll() *gorm.DB {
	items := []Product{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB {
	var item Product
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Product) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates(id string, newProduct *Product) *gorm.DB {
	var item Product
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func Delete(id string) *gorm.DB {
	var item Product
	return config.DB.Delete(&item, "id = ?", id)
}

func FindData(product_name string) *gorm.DB {
	items := []Product{}
	product_name = "%" + product_name + "%"
	return config.DB.Where("product_name LIKE ?", product_name).Find(&items)
}

func FindCond(sort string, limit int, offset int) *gorm.DB {
	items := []Product{}
	return config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
}

func CountData() int {
	var result int
	config.DB.Table("products").Count(&result)
	return result
}
