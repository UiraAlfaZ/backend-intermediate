package helpers

import (
	"Backend/backend-api/src/config"
	"Backend/backend-api/src/models/customer"
	products "Backend/backend-api/src/models/products"
	"Backend/backend-api/src/models/seller"
	"Backend/backend-api/src/models/transaction"
	models "Backend/backend-api/src/models/user"
)

func Migration() {
	config.DB.AutoMigrate(&customer.Customer{})
	config.DB.AutoMigrate(&seller.Seller{})
	config.DB.AutoMigrate(&transaction.Transaction{})
	config.DB.AutoMigrate(&products.Product{})
	config.DB.AutoMigrate(&models.User{})

}
