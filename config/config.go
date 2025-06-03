package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	TimeExpiredAt = time.Hour * 24
)

type Config struct {
	Environment string

	ServerHost string
	ServerPort string


	Postgres Postgres

	SekretKey string
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DataBase string
}

func Load() *Config {
	if err := godotenv.Load("config/.env"); err != nil {
		fmt.Println("NO .env file  not found")
	}

	cfg := Config{}
	cfg.ServerHost = cast.ToString(getOrDefaultValue("SERVER_HOST", "62.171.149.94"))
	cfg.Environment = cast.ToString(getOrDefaultValue("ENVIRONMENT", "dev"))
	cfg.Postgres = Postgres{
		Host:     cast.ToString(getOrDefaultValue("POSTGRES_HOST", "62.171.149.94")),
		Port:     cast.ToInt(getOrDefaultValue("POSTGRES_PORT", "5432")),
		User:     cast.ToString(getOrDefaultValue("POSTGRES_USER", "maxmurjon")),
		Password: cast.ToString(getOrDefaultValue("POSTGRES_PASSWORD", "max22012004")),
		DataBase: cast.ToString(getOrDefaultValue("POSTGRES_DATABASE", "potgres"))}

	cfg.SekretKey = cast.ToString(getOrDefaultValue("JWT_SECRET", "sekret"))
	return &cfg
}

func getOrDefaultValue(key string, defaultValue string) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
