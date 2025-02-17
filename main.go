package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/ssssshel/sp-api/src/config"
	"github.com/ssssshel/sp-api/src/middlewares"
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
		fmt.Println("Error loading .env file: ", err)
	}
}

// createServer creates a new http.Server and returns a pointer to it
func createServer(tokenization bool) *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	var handler http.Handler = mux
	handler = middlewares.Cors(handler)
	handler = middlewares.Logger(handler)

	if tokenization {
		println("Tokenization is enabled")
	}

	server := &http.Server{
		Addr:    ":" + config.Env().Port,
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
		fmt.Println("Running in development/testing mode")
	case Production:
		fmt.Println("Running in production mode")
	}

	tokenization := environment != DevelopmentWithoutTokens
	server := createServer(tokenization)

	fmt.Printf("Server running on port %s\n", config.Env().Port)
	log.Fatal(server.ListenAndServe())
}

func main() {
	loadEnv()
	startServer(Env(config.Env().EnvCode))
}
