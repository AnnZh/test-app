package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	test_app "github.com/AnnZh/test-app"
	"github.com/AnnZh/test-app/pkg/handler"
	"github.com/AnnZh/test-app/pkg/repository"
	"github.com/AnnZh/test-app/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	f, err := repository.NewFile(viper.GetString("file.name"))
	if err != nil {
		logrus.Fatalf("failed to initialize file %s", err.Error())
	}

	repos := repository.NewRepository(f)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(test_app.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Test App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Test App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := f.Close(); err != nil {
		logrus.Errorf("error occured on file connection close: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
