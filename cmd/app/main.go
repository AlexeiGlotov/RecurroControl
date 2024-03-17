package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	todo "RecurroControl"
	"RecurroControl/internal/handler"
	repository2 "RecurroControl/internal/repository"
	"RecurroControl/internal/service"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error configs %s", err.Error())
	}

	db, err := repository2.NewMysqlDB(repository2.Config{
		Host:     os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     viper.GetString("db.port"),
		Login:    viper.GetString("db.login"),
		DBName:   viper.GetString("db.dbname"),
	})

	if err != nil {
		logrus.Fatalf("error mysql %s", err.Error())
	}

	repos := repository2.NewRepository(db)
	services := service.NewService(repos, os.Getenv("SALT_PASSWORD"), os.Getenv("SALT_JWT"))
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
