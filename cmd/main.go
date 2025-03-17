package main

import (
	restapigarbage "github.com/Xu3is/RestApiGarbage"
	"github.com/Xu3is/RestApiGarbage/pkg/handler"
	"github.com/Xu3is/RestApiGarbage/pkg/repository"
	"github.com/Xu3is/RestApiGarbage/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
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

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file, %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Error connecting to database, %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restapigarbage.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while starting http server: %v", err.Error())
	}
}
