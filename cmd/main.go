package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sinyakovskiy/todo-app"
	"github.com/sinyakovskiy/todo-app/pkg/handler"
	"github.com/sinyakovskiy/todo-app/pkg/repository"
	"github.com/sinyakovskiy/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:        viper.GetString("db.host"),
		Port:        viper.GetString("db.port"),
		Username:    viper.GetString("db.username"),
		Password:    os.Getenv("DB_PASSWORD"),
		DBName:      viper.GetString("db.dbname"),
		SSLMode:     viper.GetString("db.sslmode"),
		SSLRootCert: viper.GetString("db.sslrootcert"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHundler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
