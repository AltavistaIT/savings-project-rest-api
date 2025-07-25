package router_user

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	handler_user "github.com/ssssshel/sp-api/src/presentation/http/handler/User"
	"github.com/ssssshel/sp-api/src/shared"
	usecases_user "github.com/ssssshel/sp-api/src/usecases/User"
)

func UserRoutes(mux *http.ServeMux, container *shared.Container) {
	userRepository := infra_db.NewUserRepository(container.DB.DBConn)
	getUserByIdUsecase := usecases_user.NewGetUserByIdUsecase(userRepository)
	userHandler := handler_user.NewUserHandler(getUserByIdUsecase)

	mux.HandleFunc("GET /users/", func(w http.ResponseWriter, r *http.Request) {
		userHandler.GetById(w, r)
	})
}
