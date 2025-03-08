package routes

import (
	"net/http"

	"github.com/ssssshel/sp-api/src/controllers"
)

func UserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetUsers(w, r)
		case http.MethodPost:
			controllers.CreateUser(w, r)
		}
	})
}
