package router_user

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	handler_user "github.com/ssssshel/sp-api/src/presentation/http/handler/User"
	"github.com/ssssshel/sp-api/src/shared"
	"github.com/ssssshel/sp-api/src/shared/logger"
	usecases_user "github.com/ssssshel/sp-api/src/usecases/User"
)

func UserRoutes(mux *http.ServeMux, container *shared.Container) {
	logger.Info("Configuring user routes")
	userRepository := infra_db.NewUserRepository(container.DB.DBConn)
	createUserUsecase := usecases_user.NewCreateUserUsecase(userRepository)
	getUserByIdUsecase := usecases_user.NewGetUserByIdUsecase(userRepository)
	userHandler := handler_user.NewUserHandler(createUserUsecase, getUserByIdUsecase)

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.Create(w, r)
		default:
			handler.HandleHttpError(w, http.StatusMethodNotAllowed, nil)
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetById(w, r)
		default:
			handler.HandleHttpError(w, http.StatusMethodNotAllowed, nil)
		}
	})
}
