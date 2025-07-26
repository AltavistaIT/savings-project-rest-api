package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	"github.com/ssssshel/sp-api/src/shared/config"
	"github.com/ssssshel/sp-api/src/shared/logger"
	usecases_auth "github.com/ssssshel/sp-api/src/usecases/Auth"
)

type AuthorizationMiddleware struct {
	jwtUsecase usecases_auth.JWTUsecase
}

func NewAuthorizationMiddleware() *AuthorizationMiddleware {
	return &AuthorizationMiddleware{
		jwtUsecase: usecases_auth.NewJWTUsecase(config.GetConfig().JWTSecretKey, config.GetConfig().JWTIssuer),
	}
}

func (m *AuthorizationMiddleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handler.HandleHttpError(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			handler.HandleHttpError(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		tokenString := parts[1]
		userID, err := m.jwtUsecase.ValidateToken(tokenString)
		logger.Info("err: %v", err)
		if err != nil {
			handler.HandleHttpError(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
