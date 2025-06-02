package router

import (
	"net/http"

	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	router_table "github.com/ssssshel/sp-api/src/presentation/http/router/Table"
	router_transaction "github.com/ssssshel/sp-api/src/presentation/http/router/Transaction"
	router_user "github.com/ssssshel/sp-api/src/presentation/http/router/User"
	"github.com/ssssshel/sp-api/src/shared"
)

func InitRoutes(container *shared.Container) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is up and running"))
	})

	apiMux := http.NewServeMux()
	router_table.TableRouter(apiMux, container)
	router_transaction.TransactionRoutes(apiMux, container)
	router_user.UserRoutes(apiMux, container)

	mux.Handle("/api/", http.StripPrefix("/api/", apiMux))

	// Handle not found
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleHttpError(w, http.StatusNotFound, "Not found", nil)
	})

	return mux
}
