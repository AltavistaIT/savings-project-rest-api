package controllers

import "net/http"

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a user"))
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}
