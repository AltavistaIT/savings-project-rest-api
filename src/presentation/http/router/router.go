package router

import (
	"net/http"

	"github.com/ssssshel/sp-api/src/presentation/http/handler"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is up and running"))
	})

	apiMux := http.NewServeMux()
	mux.Handle("/api/", apiMux)

	mux.Handle("/api/", http.StripPrefix("/api/", apiMux))

	// Handle not found
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleHttpError(w, http.StatusNotFound, "Not found", nil)
	})

	return mux
}
