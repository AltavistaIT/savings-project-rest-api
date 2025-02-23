package controllers

import "net/http"

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all transactions"))
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a transaction"))
}
