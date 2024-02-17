package controllers_products

import (
	"Backend/backend-api/src/helpers"
	"Backend/backend-api/src/middleware"
	products "Backend/backend-api/src/models/products"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Data_products(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helpers.EnableCors(w)
	if r.Method == "GET" {
		pageOld := r.URL.Query().Get("page")
		limitOld := r.URL.Query().Get("limit")
		page, _ := strconv.Atoi(pageOld)
		limit, _ := strconv.Atoi(limitOld)
		offset := (page - 1) * limit
		sort := r.URL.Query().Get("sort")
		if sort == "" {
			sort = "ASC"
		}
		sortby := r.URL.Query().Get("sortBy")
		if sortby == "" {
			sortby = "name"
		}
		sort = sortby + " " + strings.ToLower(sort)
		respons := products.FindCond(sort, limit, offset)
		totalData := products.CountData()
		totalPage := math.Ceil(float64(totalData) / float64(limit))
		result := map[string]interface{}{
			"status":      "Berhasil",
			"data":        respons.Value,
			"currentPage": page,
			"limit":       limit,
			"totalData":   totalData,
			"totalPage":   totalPage,
		}
		res, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if r.Method == "POST" {
		var product products.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode")
			return
		}

		products.Post(&product)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Product Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_product(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helpers.EnableCors(w)
	id := r.URL.Path[len("/products/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(products.Select(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if r.Method == "PUT" {
		var updateProduct products.Product
		err := json.NewDecoder(r.Body).Decode(&updateProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode boss")
			return
		}

		NewProduct := products.Product{
			IdItem:         updateProduct.IdItem,
			ProductName:    updateProduct.ProductName,
			DateProduction: updateProduct.DateProduction,
			Condition:      updateProduct.Condition,
			Stock:          updateProduct.Stock,
		}

		products.Updates(id, &NewProduct)
		msg := map[string]string{
			"Message": "Product Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if r.Method == "DELETE" {
		products.Delete(id)
		msg := map[string]string{
			"Message": "Product Deleted",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

// Upload//
func Handle_upload(w http.ResponseWriter, r *http.Request) {
	const (
		AllowedExtensions = ".jpg,.jpeg,.pdf,.png"
		MaxFileSize       = 2 << 20
	)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	ext := filepath.Ext(handler.Filename)
	ext = strings.ToLower(ext)
	allowedExts := strings.Split(AllowedExtensions, ",")
	validExtension := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			validExtension = true
			break
		}
	}
	if !validExtension {
		http.Error(w, "Invalid file extension", http.StatusBadRequest)
		return
	}

	fileSize := handler.Size
	if fileSize > MaxFileSize {
		http.Error(w, "File size exceeds the allowed limit", http.StatusBadRequest)
		return
	}

	timeStamp := time.Now().Format("20060102_150405")

	filename := fmt.Sprintf("src/upload/%s_%s", timeStamp, handler.Filename)

	out, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := map[string]string{
		"Message": "File uploaded succesfully",
	}
	res, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		return
	}

	w.Write(res)

}

// Search//
func SearchProduct(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("search")
	res, err := json.Marshal(products.FindData(keyword).Value)
	if err != nil {
		http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		return
	}
	w.Write(res)

}

// item := products.Product{
// 	IdItem:         input.IdItem,
// 	ProductName:    input.ProductName,
// 	DateProduction: input.DateProduction,
// 	Condition:      input.Condition,
// 	Stock:          input.Stock,
// }

// func Data_products(w http.ResponseWriter, r *http.Request) {
// 	middleware.GetCleanedInput(r)
// 	helpers.EnableCors(w)
// 	if r.Method == "GET" {
// 		res, err := json.Marshal(products.SelectAll().Value)
// 		if err != nil {
// 			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(res)
// 		w.Header().Set("Content-Type", "application/json")
// 		return
// 	} else if r.Method == "POST" {
// 		var input products.Product
// 		err := json.NewDecoder(r.Body).Decode(&input)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		Product := products.Product{
// 			IdItem:      input.IdItem,
// 			ProductName: input.ProductName,
// 			Stock:       input.Stock,
// 		}
// 		products.Post(&Product)
// 		w.WriteHeader(http.StatusCreated)
// 		msg := map[string]string{
// 			"Message": "Id Created",
// 		}
// 		res, err := json.Marshal(msg)
// 		if err != nil {
// 			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(res)
// 	} else {
// 		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
// 	}
// }

// func Data_product(w http.ResponseWriter, r *http.Request) {
// 	middleware.GetCleanedInput(r)
// 	helpers.EnableCors(w)
// 	id := r.URL.Path[len("/product/"):]

// 	if r.Method == "GET" {
// 		res, err := json.Marshal(products.Select(id).Value)
// 		if err != nil {
// 			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
// 		}
// 		w.Write(res)
// 		w.Header().Set("Content-Type", "application/json")
// 		return
// 	} else if r.Method == "PUT" {
// 		var updateItem products.Product
// 		err := json.NewDecoder(r.Body).Decode(&updateItem)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		NewItem := products.Product{
// 			IdItem:      updateItem.IdItem,
// 			ProductName: updateItem.ProductName,
// 			Stock:       updateItem.Stock,
// 		}
// 		products.Updates(id, &NewItem)
// 		msg := map[string]string{
// 			"Message": "Data Updated",
// 		}
// 		res, err := json.Marshal(msg)
// 		if err != nil {
// 			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(res)
// 	} else if r.Method == "DELETE" {
// 		products.Deletes(id)
// 		msg := map[string]string{
// 			"Message": "Product Deleted",
// 		}
// 		res, err := json.Marshal(msg)
// 		if err != nil {
// 			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(res)
// 	} else {
// 		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
// 	}
// }
