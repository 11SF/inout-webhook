package main

import (
	"fmt"
	"os"

	"github.com/11SF/inout-webhook/configs"
	routers "github.com/11SF/inout-webhook/pkg"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {

	slog.Info("[server] starting")
	config := configs.LoadEnv()

	server := routers.NewRouters(config).InitRouters()
	startServer(server, config)
}

func startServer(server *gin.Engine, config *configs.Config) {
	server.Run(fmt.Sprintf(":%v", config.AppPort))
}
