package routes

import (
	"net/http"

	"github.com/ssssshel/sp-api/src/controllers"
)

func TransactionRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/transactions", controllers.GetTransactions)
	mux.HandleFunc("/transactions/", controllers.GetTransaction)
}
