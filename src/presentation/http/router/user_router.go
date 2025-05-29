package routes

import (
	"errors"
	"net/http"

	controller "github.com/ssssshel/sp-api/src/controllers/users"
	"github.com/ssssshel/sp-api/src/utils"
)

func UserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.CreateUser(w, r)
		default:
			utils.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("method not allowed"))
		}
	})

	mux.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.GetUser(w, r)
		default:
			utils.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("method not allowed"))
		}
	})
}
