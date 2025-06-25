package handler_config

import (
	"net/http"

	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	usecases_config "github.com/ssssshel/sp-api/src/usecases/Config"
)

type ConfigHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type configHandler struct {
	configUsecase usecases_config.GetConfigUsecase
}

func NewConfigHandler(configUsecase usecases_config.GetConfigUsecase) ConfigHandler {
	return &configHandler{
		configUsecase: configUsecase,
	}
}

func (ch *configHandler) Get(w http.ResponseWriter, r *http.Request) {
	config, err := ch.configUsecase.Execute()

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusOK, handler.HttpMessage[http.StatusOK], config)
}
