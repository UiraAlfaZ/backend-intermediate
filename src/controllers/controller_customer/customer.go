package controllers_customer

import (
	"Backend/backend-api/src/helpers"
	"Backend/backend-api/src/middleware"
	"Backend/backend-api/src/models/customer"
	"encoding/json"
	"net/http"
)

func Data_customers(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helpers.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(customer.SelectAllOrder().Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		return
	} else if r.Method == "POST" {
		var input customer.Customer
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		profileCustomer := customer.Customer{
			IdCustomer: input.IdCustomer,
			Name:       input.Name,
			Address:    input.Address,
		}
		customer.Order(&profileCustomer)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Your order has been sucessful",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_customer(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helpers.EnableCors(w)
	id := r.URL.Path[len("/customer/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(customer.Select(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		return
	} else if r.Method == "PUT" {
		var updateCustomer customer.Customer
		err := json.NewDecoder(r.Body).Decode(&updateCustomer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newCustomer := customer.Customer{
			IdCustomer: updateCustomer.IdCustomer,
			Name:       updateCustomer.Name,
			Address:    updateCustomer.Address,
		}
		customer.Updates(id, &newCustomer)
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
		customer.Deletes(id)
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
