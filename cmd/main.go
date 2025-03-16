package main

import (
	"context"
	todo "github.com/Xu3is/RestApiGarbage"
	"github.com/Xu3is/RestApiGarbage/pkg/handler"
	"github.com/Xu3is/RestApiGarbage/pkg/repository"
	"github.com/Xu3is/RestApiGarbage/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func initConfig() error {
	viper.AddConfigPath("cfgs")
	viper.SetConfigName("cfg")
	return viper.ReadInConfig()
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("can't init config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initializa db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("failed to initialize server: %s", err.Error())
		}
	}()
	logrus.Print("Application started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Application stopped")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("error occure on db connection close %s", err.Error())
	}

}
