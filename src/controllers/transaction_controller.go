package controllers

import "net/http"

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all transactions"))
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a transaction"))
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a transaction"))
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a transaction"))
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a transaction"))
}
