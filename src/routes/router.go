package routes

import (
	"net/http"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("Welcome to the API"))
	})
	TransactionRoutes(mux)
	UserRoutes(mux)

	return mux
}
