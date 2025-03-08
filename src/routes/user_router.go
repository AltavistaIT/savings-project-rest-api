package routes

import (
	"errors"
	"net/http"

	"github.com/ssssshel/sp-api/src/controllers"
	"github.com/ssssshel/sp-api/src/utils"
)

func UserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetUsers(w, r)
		case http.MethodPost:
			controllers.CreateUser(w, r)
		default:
			utils.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("method not allowed"))
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetUser(w, r)
		default:
			utils.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("method not allowed"))
		}
	})
}
