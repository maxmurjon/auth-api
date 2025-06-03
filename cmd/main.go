package main

import (
	"fmt"
	"smartlogistics/api"
	"smartlogistics/api/handler"
	"smartlogistics/config"

	postgres "smartlogistics/storage/postges"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	psqlConnString := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.DataBase,
		cfg.Postgres.Password,
		cfg.Postgres.Port,
	)

	strg := postgres.NewPostgres(psqlConnString)

	h := handler.NewHandler(cfg, strg)

	switch cfg.Environment {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	api.SetUpAPI(r, *h, *cfg)

	fmt.Println("Server running on port 8000")
	r.Run(":8000")
}
