package routes

import (
	"net/http"

	"github.com/ssssshel/sp-api/src/controllers"
)

func TransactionRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetTransactions(w, r)
		case http.MethodPost:
			controllers.CreateTransaction(w, r)
		}
	})

	mux.HandleFunc("/transactions/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetTransaction(w, r)
		case http.MethodPatch:
			controllers.UpdateTransaction(w, r)
		case http.MethodDelete:
			controllers.DeleteTransaction(w, r)
		}
	})
}
