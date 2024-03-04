package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"

	todo "RecurroControl"
	"RecurroControl/pkg/handler"
	"RecurroControl/pkg/repository"
	"RecurroControl/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error configs %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error env %s", err.Error())
	}

	db, err := repository.NewMysqlDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     viper.GetString("db.port"),
		Login:    viper.GetString("db.login"),
		DBName:   viper.GetString("db.dbname"),
	})

	if err != nil {
		logrus.Fatalf("error mysql %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
