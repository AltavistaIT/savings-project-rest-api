package router_transaction

import (
	"errors"
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	handler_transaction "github.com/ssssshel/sp-api/src/presentation/http/handler/Transaction"
	"github.com/ssssshel/sp-api/src/shared"
	usecases_transaction "github.com/ssssshel/sp-api/src/usecases/Transaction"
)

func TransactionRoutes(mux *http.ServeMux, container *shared.Container) {
	transacionRepository := infra_db.NewTransactionRepository(container.DB.DBConn)
	transactionTableRepository := infra_db.NewTransactionTableRepository(container.DB.DBConn)
	getTransactionsByTableId := usecases_transaction.NewGetTransactionsByTableIdUsecase(transacionRepository)
	createTransactionUsecase := usecases_transaction.NewCreateTransactionUsecase(transacionRepository, transactionTableRepository)
	transactionHandler := handler_transaction.NewTransactionHandler(createTransactionUsecase, getTransactionsByTableId)

	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			transactionHandler.Create(w, r)
		default:
			handler.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("method not allowed"))
		}
	})

	mux.HandleFunc("/transactions/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			transactionHandler.GetByTableId(w, r)
		default:
			handler.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("method not allowed"))
		}
	})
}
