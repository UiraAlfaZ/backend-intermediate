package controllers_seller

import (
	"Backend/backend-api/src/helpers"
	"Backend/backend-api/src/middleware"
	"Backend/backend-api/src/models/seller"

	"encoding/json"
	"net/http"
)

func Data_sellers(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helpers.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(seller.SelectAll().Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		return
	} else if r.Method == "POST" {
		var input seller.Seller
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		profileSeller := seller.Seller{
			IdSeller: input.IdSeller,
			Name:     input.Name,
		}
		seller.Post(&profileSeller)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Id Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_seller(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helpers.EnableCors(w)
	id := r.URL.Path[len("/seller/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(seller.Select(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		return
	} else if r.Method == "PUT" {
		var updateSeller seller.Seller
		err := json.NewDecoder(r.Body).Decode(&updateSeller)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newSeller := seller.Seller{
			IdSeller: updateSeller.IdSeller,
			Name:     updateSeller.Name,
		}
		seller.Updates(id, &newSeller)
		msg := map[string]string{
			"Message": "Data Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else if r.Method == "DELETE" {
		seller.Deletes(id)
		msg := map[string]string{
			"Message": "Product Deleted",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
