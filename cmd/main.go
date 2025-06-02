package main

import (
	"github.com/ssssshel/sp-api/cmd/server"
	"github.com/ssssshel/sp-api/src/shared/config"
)

func main() {
	config := config.GetConfig()
	server := server.NewServer(config)
	server.Start()
}
