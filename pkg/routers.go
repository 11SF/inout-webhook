package routers

import (
	"net/http"

	"github.com/11SF/inout-webhook/configs"
	httpinout "github.com/11SF/inout-webhook/pkg/v1/core/http"
	coreline "github.com/11SF/inout-webhook/pkg/v1/core/line_event"
	linehandleer "github.com/11SF/inout-webhook/pkg/v1/line"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Routers struct {
	config    *configs.Config
	lineBot   *linebot.Client
	httoInOut *httpinout.HTTPInOut
}

func NewRouters(config *configs.Config, lineBot *linebot.Client, httoInOut *httpinout.HTTPInOut) *Routers {
	return &Routers{config, lineBot, httoInOut}
}

func (router *Routers) InitRouters() *gin.Engine {
	if router.config.AppEnviroment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	contextPath := r.Group("/inout-webhook")
	v1 := contextPath.Group("/v1")

	lineService := coreline.NewService(router.httoInOut, router.config)
	lineHandler := linehandleer.NewLineHandler(router.lineBot, lineService.EventMessage)
	v1.POST("/line", lineHandler.LineHandler)

	return r
}
