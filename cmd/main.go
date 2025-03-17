package main

import (
	restapigarbage "github.com/Xu3is/RestApiGarbage"
	"github.com/Xu3is/RestApiGarbage/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(restapigarbage.Server)
	err := srv.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while starting http server: %v", err.Error())
	}
}
