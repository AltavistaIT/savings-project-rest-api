package router

import (
	"net/http"

	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	router_auth "github.com/ssssshel/sp-api/src/presentation/http/router/Auth"
	router_config "github.com/ssssshel/sp-api/src/presentation/http/router/Config"
	router_table "github.com/ssssshel/sp-api/src/presentation/http/router/Table"
	router_transaction "github.com/ssssshel/sp-api/src/presentation/http/router/Transaction"
	"github.com/ssssshel/sp-api/src/shared"
)

func InitRoutes(container *shared.Container) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is up and running"))
	})

	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./openapi.yaml")
	})

	apiMux := http.NewServeMux()
	router_auth.AuthRouter(apiMux, container)
	router_config.ConfigRouter(apiMux, container)
	router_table.TableRouter(apiMux, container)
	router_transaction.TransactionRoutes(apiMux, container)

	mux.Handle("/api/", http.StripPrefix("/api", apiMux))

	// Handle not found
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleHttpError(w, http.StatusNotFound, nil)
	})

	return mux
}
