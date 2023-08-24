package main

import (
	"fmt"
	"os"

	"github.com/11SF/inout-webhook/configs"
	routers "github.com/11SF/inout-webhook/pkg"
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

	server := routers.NewRouters(config, bot).InitRouters()
	startServer(server, config)
}

func startServer(server *gin.Engine, config *configs.Config) {
	server.Run(fmt.Sprintf(":%v", config.AppPort))
}
