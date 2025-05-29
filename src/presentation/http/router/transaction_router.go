package routes

import (
	"errors"
	"net/http"

	controller "github.com/ssssshel/sp-api/src/controllers/transactions"
	"github.com/ssssshel/sp-api/src/utils"
)

func TransactionRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// case http.MethodGet:
		// 	controllers.GetTransactions(w, r)
		case http.MethodPost:
			controller.CreateTransaction(w, r)
		default:
			utils.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("method not allowed"))
		}
	})

	// mux.HandleFunc("/api/transactions/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		controllers.GetTransaction(w, r)
	// 	case http.MethodPatch:
	// 		controllers.UpdateTransaction(w, r)
	// 	case http.MethodDelete:
	// 		controllers.DeleteTransaction(w, r)
	// 	}
	// })
}
