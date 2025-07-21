package router_transaction

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	handler_transaction "github.com/ssssshel/sp-api/src/presentation/http/handler/Transaction"
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

	mux.HandleFunc("POST /transactions", func(w http.ResponseWriter, r *http.Request) {
		transactionHandler.Create(w, r)
	})

	mux.HandleFunc("PATCH /transactions/{id}", func(w http.ResponseWriter, r *http.Request) {
		transactionHandler.Update(w, r)
	})

	mux.HandleFunc("DELETE /transactions/{id}", func(w http.ResponseWriter, r *http.Request) {
		transactionHandler.Delete(w, r)
	})
}
