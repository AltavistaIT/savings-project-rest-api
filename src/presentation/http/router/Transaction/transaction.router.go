package router_transaction

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	handler_transaction "github.com/ssssshel/sp-api/src/presentation/http/handler/Transaction"
	middlewares "github.com/ssssshel/sp-api/src/presentation/http/middleware"
	"github.com/ssssshel/sp-api/src/shared"
	usecases_transaction "github.com/ssssshel/sp-api/src/usecases/Transaction"
)

func TransactionRoutes(mux *http.ServeMux, container *shared.Container) {
	transacionRepository := infra_db.NewTransactionRepository(container.DB.DBConn)
	transactionTableRepository := infra_db.NewTransactionTableRepository(container.DB.DBConn)
	tableRepository := infra_db.NewTableRepository(container.DB.DBConn)
	createTransactionUsecase := usecases_transaction.NewCreateTransactionUsecase(transacionRepository, transactionTableRepository, tableRepository)
	updateTransactionUsecase := usecases_transaction.NewUpdateTransactionUsease(transacionRepository)
	deleteTransactionUscase := usecases_transaction.NewDeleteTransactionUsecase(transacionRepository, transactionTableRepository)
	transactionHandler := handler_transaction.NewTransactionHandler(createTransactionUsecase, updateTransactionUsecase, deleteTransactionUscase)

	auth := middlewares.NewAuthorizationMiddleware().Authorization

	mux.Handle("POST /transactions", shared.Protect(auth, transactionHandler.Create))

	mux.Handle("PATCH /transactions/{id}", shared.Protect(auth, transactionHandler.Update))

	mux.Handle("DELETE /transactions/{id}", shared.Protect(auth, transactionHandler.Delete))
}
