package router_table

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	handler_table "github.com/ssssshel/sp-api/src/presentation/http/handler/Table"
	"github.com/ssssshel/sp-api/src/shared"
	usecases_table "github.com/ssssshel/sp-api/src/usecases/Table"
)

func TableRouter(mux *http.ServeMux, container *shared.Container) {
	tableRepository := infra_db.NewTableRepository(container.DB.DBConn)
	transactionRepository := infra_db.NewTransactionRepository(container.DB.DBConn)
	getTableByIdUsecase := usecases_table.NewGetTableByIdUsecase(tableRepository, transactionRepository)
	createTableUsecase := usecases_table.NewCreateTableUsecase(tableRepository)
	tableHandler := handler_table.NewTableHandler(createTableUsecase, getTableByIdUsecase)

	mux.HandleFunc("/tables", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			tableHandler.Create(w, r)
		default:
			handler.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

	mux.HandleFunc("/tables/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tableHandler.GetById(w, r)
		default:
			handler.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})
}
