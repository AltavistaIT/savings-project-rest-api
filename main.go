package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/ssssshel/sp-api/src/config"
	"github.com/ssssshel/sp-api/src/middlewares"
	"github.com/ssssshel/sp-api/src/routes"
)

type Env uint

const (
	Development              Env = 1 // DO NOT USE THIS IN PRODUCTION
	Testing                  Env = 2 // DO NOT USE THIS IN PRODUCTION
	DevelopmentWithoutTokens Env = 3 // DO NOT USE THIS IN PRODUCTION
	Production               Env = 4
)

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file: ", err)
	}
}

// createServer creates a new http.Server and returns a pointer to it
func createServer(tokenization bool) *http.Server {
	config.DBConnection()
	// config.MigrateDB()

	mux := routes.InitRoutes()

	var handler http.Handler = mux
	handler = middlewares.Cors(handler)
	handler = middlewares.Logger(handler)

	if tokenization {
		log.Println("Tokenization is enabled")
	}

	server := &http.Server{
		Addr:    ":" + config.ServiceConfig().Port,
		Handler: handler,
	}
	return server
}

// startServer starts the server based on the environment configuration
func startServer(environment Env) {
	if environment < Development || environment > Production {
		log.Fatal("Invalid environment configuration")
	}

	switch environment {
	case Development, Testing, DevelopmentWithoutTokens:
		log.Println("Running in development/testing mode")
	case Production:
		log.Println("Running in production mode")
	}

	tokenization := environment != DevelopmentWithoutTokens
	server := createServer(tokenization)

	log.Printf("Server running on port %s\n", config.ServiceConfig().Port)
	log.Fatal(server.ListenAndServe())
}

func main() {
	loadEnv()
	startServer(Env(config.ServiceConfig().EnvCode))
}
