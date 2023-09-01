package main

import (
	"github.com/Andrew-2609/go-opportunities/config"
	"github.com/Andrew-2609/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v\n", err)
		return
	}

	router.Initialize()
}
