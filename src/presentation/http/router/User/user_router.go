package router_user

import (
	"errors"
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	handler_user "github.com/ssssshel/sp-api/src/presentation/http/handler/User"
	"github.com/ssssshel/sp-api/src/shared"
	usecases_user "github.com/ssssshel/sp-api/src/usecases/User"
)

func UserRoutes(mux *http.ServeMux, container *shared.Container) {
	userRepository := infra_db.NewUserRepository(container.DB.DBConn)
	createUserUsecase := usecases_user.NewCreateUserUsecase(userRepository)
	getUserByIdUsecase := usecases_user.NewGetUserByIdUsecase(userRepository)
	userHandler := handler_user.NewUserHandler(createUserUsecase, getUserByIdUsecase)

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.Create(w, r)
		default:
			handler.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("method not allowed"))
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetById(w, r)
		default:
			handler.HandleHttpError(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("method not allowed"))
		}
	})
}
