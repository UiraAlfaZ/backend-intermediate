package routes

import (
	controllers_customer "Backend/backend-api/src/controllers/controller_customer"
	controllers_products "Backend/backend-api/src/controllers/controller_product"
	controllers_seller "Backend/backend-api/src/controllers/controller_seller"
	controllers_transaction "Backend/backend-api/src/controllers/controller_transaction"
	controllers "Backend/backend-api/src/controllers/user"
	"Backend/backend-api/src/middleware"
	"fmt"
	"net/http"

	"github.com/goddtriffin/helmet"
)

func Routes() {
	helmet := helmet.Default()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	// http.HandleFunc("/seller", controllers_seller.Data_sellers)
	// http.HandleFunc("/seller/", controllers_seller.Data_seller)
	// http.HandleFunc("/customer", controllers_customer.Data_customers)
	// http.HandleFunc("/customer/", controllers_customer.Data_customer)
	// http.HandleFunc("/transaction", controllers_transaction.Data_transactions)
	// http.HandleFunc("/transaction/", controllers_transaction.Data_transaction)
	// http.HandleFunc("/products", controllers_products.Data_products)
	// http.HandleFunc("/product/", controllers_products.Data_Product)

	//-----------------Seller Section------------------//
	http.Handle("/seller", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(controllers_seller.Data_sellers))))
	http.Handle("/seller/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(controllers_seller.Data_seller))))

	//-----------------Customer Section----------------//
	http.Handle("/customer", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(controllers_customer.Data_customers))))
	http.Handle("/customer/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(controllers_customer.Data_customer))))

	//----------------Transaction Section--------------//
	http.Handle("/transaction", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(controllers_transaction.Data_transactions))))
	http.Handle("/transaction/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(controllers_transaction.Data_transaction))))

	//----------------Product Section------------------//
	http.Handle("/products", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(controllers_products.Data_products))))
	http.Handle("/products/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(controllers_products.Data_product))))
	//----------------Search and Upload Products-------//
	http.Handle("/search", http.HandlerFunc(controllers_products.SearchProduct))
	http.Handle("/upload", http.HandlerFunc(controllers_products.Handle_upload))

	//----------------User Section---------------------//
	//////////////////register//////////////////////
	http.Handle("/register", http.HandlerFunc(controllers.RegisterUser))
	/////////////////login//////////////////////////
	http.Handle("/login", http.HandlerFunc(controllers.LoginUser))
}
