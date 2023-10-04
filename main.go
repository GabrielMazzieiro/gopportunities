package main

import (
	"fmt"

	"github.com/GabrielMazzieiro/gopportunities/config"
	"github.com/GabrielMazzieiro/gopportunities/router"
)

func main() {

	//Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	//Initialize Router
	router.Initialize()
}
