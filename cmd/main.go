package main

import (
	restapigarbage "github.com/Xu3is/RestApiGarbage"
	"github.com/Xu3is/RestApiGarbage/pkg/handler"
	"github.com/Xu3is/RestApiGarbage/pkg/repository"
	"github.com/Xu3is/RestApiGarbage/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func configInit() error {
	viper.AddConfigPath("cfgs")
	viper.SetConfigName("cfg")
	return viper.ReadInConfig()
}

func main() {
	err := configInit()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restapigarbage.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while starting http server: %v", err.Error())
	}
}
