package routers

import (
	"github.com/11SF/inout-webhook/configs"
	httpinout "github.com/11SF/inout-webhook/pkg/v1/core/http"
	coreline "github.com/11SF/inout-webhook/pkg/v1/core/line_event"
	linehandleer "github.com/11SF/inout-webhook/pkg/v1/line"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Routers struct {
	config  *configs.Config
	lineBot *linebot.Client
}

func NewRouters(config *configs.Config, lineBot *linebot.Client) *Routers {
	return &Routers{config, lineBot}
}

func (router *Routers) InitRouters() *gin.Engine {
	if router.config.AppEnviroment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	contextPath := r.Group("/inout-webhook")
	v1 := contextPath.Group("/v1")

	httoInOut := httpinout.NewHTTP(&httpinout.HTTPConfig{
		Addr:       router.config.MgmtService.Addr,
		ApiVersion: router.config.MgmtService.ApiVersion,
		Endpoints: &httpinout.Endpoints{
			AddExpenseEndpoint: router.config.MgmtService.AddExpenseEndpoint,
			AddIncomeEndpoint:  router.config.MgmtService.AddIncomeEndpoint,
		},
	})
	lineService := coreline.NewService(httoInOut, router.config)
	lineHandler := linehandleer.NewLineHandler(router.lineBot, lineService.EventMessage)
	v1.POST("/line", lineHandler.LineHandler)

	return r
}
