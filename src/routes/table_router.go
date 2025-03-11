package routes

import (
	"errors"
	"net/http"

	controller "github.com/ssssshel/sp-api/src/controllers/tables"
	"github.com/ssssshel/sp-api/src/utils"
)

func TableRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/tables", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.CreateTable(w, r)
		default:
			utils.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("method not allowed"))
		}
	})
}
