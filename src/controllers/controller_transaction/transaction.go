package controllers_transaction

import (
	"Backend/backend-api/src/helpers"
	"Backend/backend-api/src/middleware"
	"Backend/backend-api/src/models/transaction"
	"encoding/json"
	"net/http"
)

func Data_transactions(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helpers.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(transaction.SelectAll().Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		return
	} else if r.Method == "POST" {
		var input transaction.Transaction
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		bucket := transaction.Transaction{
			IdProduct:   input.IdProduct,
			ProductName: input.ProductName,
			Payment:     input.Payment,
		}
		transaction.Post(&bucket)
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

func Data_transaction(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helpers.EnableCors(w)
	id := r.URL.Path[len("/transaction/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(transaction.Select(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		return
	} else if r.Method == "PUT" {
		var updateTransaction transaction.Transaction
		err := json.NewDecoder(r.Body).Decode(&updateTransaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newTransaction := transaction.Transaction{
			IdProduct:   updateTransaction.IdProduct,
			ProductName: updateTransaction.ProductName,
			Payment:     updateTransaction.Payment,
		}
		transaction.Updates(id, &newTransaction)
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
		transaction.Deletes(id)
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
