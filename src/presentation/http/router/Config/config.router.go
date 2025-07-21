package router_config

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	handler_config "github.com/ssssshel/sp-api/src/presentation/http/handler/Config"
	"github.com/ssssshel/sp-api/src/shared"
	usecases_config "github.com/ssssshel/sp-api/src/usecases/Config"
)

func ConfigRouter(mux *http.ServeMux, container *shared.Container) {
	currencyRepository := infra_db.NewCurrencyRepository(container.DB.DBConn)
	transactionTypeRepository := infra_db.NewTransactionTypeRepository(container.DB.DBConn)
	monthYearRepository := infra_db.NewMonthYearRepository(container.DB.DBConn)
	tableTypeRepository := infra_db.NewTableTypeRepository(container.DB.DBConn)
	configUsecase := usecases_config.NewGetConfigUsecase(currencyRepository, transactionTypeRepository, monthYearRepository, tableTypeRepository)
	configHandler := handler_config.NewConfigHandler(configUsecase)

	mux.HandleFunc("GET /config", func(w http.ResponseWriter, r *http.Request) {
		configHandler.Get(w, r)
	})
}
