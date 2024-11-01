package main

import (
	"exchangeapp/EXCHANGEAPP_BACKEND/config"
	"exchangeapp/EXCHANGEAPP_BACKEND/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.Appconfig.App.Port

	if port == "" {
		port = ":8080"
	}

	r.Run(port)
}
