package routes

import (
	"net/http"

	"github.com/ssssshel/sp-api/src/utils"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/" {
			w.Write([]byte("Welcome to the API"))
			return
		}

		utils.HandleHttpError(w, http.StatusNotFound, "Route not found", nil)
	})

	TransactionRoutes(mux)
	UserRoutes(mux)
	TableRoutes(mux)

	return mux
}
