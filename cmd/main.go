package main

import (
	hezzl "Hezzl"
	"Hezzl/pkg/handler"
	"Hezzl/pkg/repository"
	"Hezzl/pkg/service"

	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	//handlers := new(handler.Handler)

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing: %s", err.Error())
	}

	/*if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}*/

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Failed to initializate db: %s", err.Error())
	}

	redis, err := repository.NewRedisClinet(repository.ConfigRedis{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	repos := repository.NewRepository(db, redis)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(hezzl.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occuped while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
