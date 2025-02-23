package routes

import (
	"net/http"

	"github.com/ssssshel/sp-api/src/controllers"
)

func UserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users/", controllers.GetUser)
}
