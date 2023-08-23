package configs

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

type Config struct {
	AppEnviroment string `env:"APP_ENVIROMENT,required"`
	AppPort       int    `env:"APP_PORT,required"`
	AppConstants  AppConstants
}

type AppConstant string

type AppConstants struct {
}

func LoadEnv() *Config {
	slog.Info("[env] start loading env")
	err := godotenv.Load("./configs/.env")
	if err != nil {
		slog.Error("[env] unable to load .env file", "error", err)
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		slog.Error("[env] unable to parse ennvironment variables", "error", err)
		panic(0)
	}

	slog.Info("[env] loading env complete")

	return cfg
}
