package router_auth

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	handler_auth "github.com/ssssshel/sp-api/src/presentation/http/handler/Auth"
	"github.com/ssssshel/sp-api/src/shared"
	"github.com/ssssshel/sp-api/src/shared/config"
	usecases_auth "github.com/ssssshel/sp-api/src/usecases/Auth"
)

func AuthRouter(mux *http.ServeMux, container *shared.Container) {
	userRepository := infra_db.NewUserRepository(container.DB.DBConn)
	jwtUsecase := usecases_auth.NewJWTUsecase(config.GetConfig().JWTSecretKey, config.GetConfig().JWTIssuer)
	loginUsecase := usecases_auth.NewLoginUsecase(userRepository, jwtUsecase)
	registerUsecase := usecases_auth.NewRegisterUsecase(userRepository)
	authHandler := handler_auth.NewAuthHandler(loginUsecase, registerUsecase)

	mux.HandleFunc("POST /auth/login", func(w http.ResponseWriter, r *http.Request) {
		authHandler.Login(w, r)
	})

	mux.HandleFunc("POST /auth/register", func(w http.ResponseWriter, r *http.Request) {
		authHandler.Register(w, r)
	})

}
