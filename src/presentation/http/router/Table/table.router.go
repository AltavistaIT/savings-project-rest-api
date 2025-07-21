package router_table

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	handler_table "github.com/ssssshel/sp-api/src/presentation/http/handler/Table"
	"github.com/ssssshel/sp-api/src/shared"
	usecases_table "github.com/ssssshel/sp-api/src/usecases/Table"
)

func TableRouter(mux *http.ServeMux, container *shared.Container) {
	tableRepository := infra_db.NewTableRepository(container.DB.DBConn)
	transactionRepository := infra_db.NewTransactionRepository(container.DB.DBConn)
	getTableByIdUsecase := usecases_table.NewGetTableByIdUsecase(tableRepository, transactionRepository)
	getTableByParamsUsecase := usecases_table.NewGetTableByParamsUsecase(tableRepository, transactionRepository)
	createTableUsecase := usecases_table.NewCreateTableUsecase(tableRepository)
	tableHandler := handler_table.NewTableHandler(createTableUsecase, getTableByIdUsecase, getTableByParamsUsecase)

	mux.HandleFunc("POST /tables", func(w http.ResponseWriter, r *http.Request) {
		tableHandler.Create(w, r)
	})

	mux.HandleFunc("GET /tables", func(w http.ResponseWriter, r *http.Request) {
		tableHandler.GetByParams(w, r)
	})

	mux.HandleFunc("GET /tables/{id}", func(w http.ResponseWriter, r *http.Request) {
		tableHandler.GetById(w, r)
	})
}
