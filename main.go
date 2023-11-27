package main

import (
	"github.com/bantheus/API-Empregos-Golang/config"
	"github.com/bantheus/API-Empregos-Golang/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
