package main

import (
	"Backend/backend-api/src/config"
	"Backend/backend-api/src/helpers"
	"Backend/backend-api/src/routes"
	"fmt"
	"net/http"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helpers.Migration()
	defer config.DB.Close()
	routes.Routes()
	fmt.Print("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
