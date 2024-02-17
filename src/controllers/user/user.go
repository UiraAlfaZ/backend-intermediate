package controllers

import (
	"Backend/backend-api/src/helpers"
	models "Backend/backend-api/src/models/user"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid request body")
			return
		}
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashPassword)
		newUser := models.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: Password,
			Address:  input.Address,
			PhoneNum: input.PhoneNum,
		}
		models.CreateUser(&newUser)
		msg := map[string]string{
			"Message": "Register successfully",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid request body")
			return
		}
		ValidateEmail := models.FindEmail(&input)
		if len(ValidateEmail) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Email is not Found")
			return
		}
		var PasswordSecond string
		for _, user := range ValidateEmail {
			PasswordSecond = user.Password
		}
		if err := bcrypt.CompareHashAndPassword([]byte(PasswordSecond), []byte(input.Password)); err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Password not Found")
			return
		}
		jwtKey := os.Getenv("SECRETKEY")
		token, _ := helpers.GenerateToken(jwtKey, input.Email)
		item := map[string]string{
			"Email": input.Email,
			"Token": token,
		}
		res, _ := json.Marshal(item)
		w.Write(res)
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}
