package router

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	handler_table "github.com/ssssshel/sp-api/src/presentation/http/handler/Table"
	usecases_table "github.com/ssssshel/sp-api/src/usecases/Table"
)

func TableRouter(mux *http.ServeMux) {

	var db *infra_db.DBConnections
	tableRepository := infra_db.NewTableRepository(db.DBConn)
	createTableUsecase := usecases_table.NewCreateTableUsecase(tableRepository)
	tableHandler := handler_table.NewTableHandler(createTableUsecase)
	mux.HandleFunc("/tables", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			tableHandler.Create(w, r)
		default:
			handler.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})
}
