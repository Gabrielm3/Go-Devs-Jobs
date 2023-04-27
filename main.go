package main

import (
	"github.com/Gabrielm3/Go-Devs-Jobs/config"
	"github.com/Gabrielm3/Go-Devs-Jobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	//initialize router
	router.Initialize()

}
