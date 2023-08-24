package main

import (
	"fmt"
	"os"

	"github.com/11SF/inout-webhook/configs"
	routers "github.com/11SF/inout-webhook/pkg"
	httpinout "github.com/11SF/inout-webhook/pkg/v1/core/http"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/exp/slog"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {

	slog.Info("[server] starting")

	config := configs.NewConfig().InitConfig()

	bot, err := linebot.New(config.Linebot.ChannelSecret, config.Linebot.ChannelToken)
	if err != nil {
		panic(err)
	}

	httoInOut := httpinout.NewHTTP(&httpinout.HTTPConfig{
		Addr:       config.MgmtService.Addr,
		ApiVersion: config.MgmtService.ApiVersion,
		Endpoints: &httpinout.Endpoints{
			AddExpenseEndpoint: config.MgmtService.AddExpenseEndpoint,
			AddIncomeEndpoint:  config.MgmtService.AddIncomeEndpoint,
		},
	})

	server := routers.NewRouters(config, bot, httoInOut).InitRouters()
	startServer(server, config)
}

func startServer(server *gin.Engine, config *configs.Config) {
	server.Run(fmt.Sprintf(":%v", config.AppPort))
}
