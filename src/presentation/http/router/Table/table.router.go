package router_table

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	handler_table "github.com/ssssshel/sp-api/src/presentation/http/handler/Table"
	middlewares "github.com/ssssshel/sp-api/src/presentation/http/middleware"
	"github.com/ssssshel/sp-api/src/shared"
	usecases_table "github.com/ssssshel/sp-api/src/usecases/Table"
)

func TableRouter(mux *http.ServeMux, container *shared.Container) {
	tableRepository := infra_db.NewTableRepository(container.DB.DBConn)
	transactionRepository := infra_db.NewTransactionRepository(container.DB.DBConn)
	getTableByParamsUsecase := usecases_table.NewGetTableByParamsUsecase(tableRepository, transactionRepository)
	createTableUsecase := usecases_table.NewCreateTableUsecase(tableRepository)
	tableHandler := handler_table.NewTableHandler(createTableUsecase, getTableByParamsUsecase)

	auth := middlewares.NewAuthorizationMiddleware().Authorization

	mux.Handle("POST /tables", shared.Protect(auth, tableHandler.Create))

	mux.Handle("GET /tables", shared.Protect(auth, tableHandler.GetByParams))
}
