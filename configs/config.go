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
	Linebot       Linebot
	MgmtService   MgmtService
}

type AppConstant string

type AppConstants struct {
	ExpenseFlag AppConstant
	IncomeFlag  AppConstant
}

type Linebot struct {
	ChannelSecret string `env:"LINE_CHANNEL_SECRET,required"`
	ChannelToken  string `env:"LINE_CHANNEL_TOKEN,required"`
}

type MgmtService struct {
	Addr               string `env:"MGMT_SERVICE_ADDR,required"`
	ApiVersion         string `env:"MGMT_SERVICE_API_VERSION,required"`
	AddExpenseEndpoint string `env:"MGMT_SERVICE_ADD_EXPENSE_ENDPOINT,required"`
	AddIncomeEndpoint  string `env:"MGMT_SERVICE_ADD_INCOME_ENDPOINT,required"`
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) initAppConstrant() {
	cfg.AppConstants.IncomeFlag = "+"
	cfg.AppConstants.ExpenseFlag = "-"
}

func (cfg *Config) loadEnv() {
	slog.Info("[env] start loading env")
	err := godotenv.Load("./configs/.env")
	if err != nil {
		slog.Error("[env] unable to load .env file", "error", err)
	}
	err = env.Parse(cfg)
	if err != nil {
		slog.Error("[env] unable to parse ennvironment variables", "error", err)
		panic(0)
	}
	slog.Info("[env] loading env complete")
}

func (cfg *Config) InitConfig() *Config {
	cfg.initAppConstrant()
	cfg.loadEnv()
	return cfg
}
