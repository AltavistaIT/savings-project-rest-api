package server

import (
	"net/http"

	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	infra_redis "github.com/ssssshel/sp-api/src/infraestructure/redis"
	middlewares "github.com/ssssshel/sp-api/src/presentation/http/middleware"
	"github.com/ssssshel/sp-api/src/presentation/http/router"
	"github.com/ssssshel/sp-api/src/shared"
	"github.com/ssssshel/sp-api/src/shared/config"
	"github.com/ssssshel/sp-api/src/shared/logger"
)

type Server interface {
	Start()
}

type server struct {
	config *config.Config
}

func NewServer(config *config.Config) Server {
	return &server{
		config: config,
	}
}

func (s *server) Start() {
	dbConnections, err := infra_db.InitConnections()
	if err != nil {
		logger.Fatal("Error connecting to database => ", err)
	}

	redisConnection, err := infra_redis.InitRedisConnection()
	if err != nil {
		logger.Fatal("Error connecting to redis => ", err)
	}

	// migrations.MigrateDB(dbConnections)

	container := &shared.Container{
		DB:    dbConnections,
		Redis: redisConnection,
	}
	mux := router.InitRoutes(container)

	var handler http.Handler = mux
	handler = middlewares.Cors(handler)
	handler = middlewares.Logger(handler)

	serverPort := ":" + s.config.Port

	app := &http.Server{
		Addr:    serverPort,
		Handler: handler,
	}
	logger.Info("Server is running on port " + serverPort)

	err = app.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal("Error starting server => ", err)
	}
}
