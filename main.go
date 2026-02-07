package main

import (
	"fmt"

	"github.com/toffettl/gopportunities/config"
	"github.com/toffettl/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
