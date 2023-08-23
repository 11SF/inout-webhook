package routers

import (
	"github.com/11SF/inout-webhook/configs"
	coreline "github.com/11SF/inout-webhook/pkg/v1/core/line_event"
	linehandleer "github.com/11SF/inout-webhook/pkg/v1/line"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Routers struct {
	config *configs.Config
}

func NewRouters(config *configs.Config) *Routers {
	return &Routers{config}
}

func (router *Routers) InitRouters() *gin.Engine {
	if router.config.AppEnviroment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	contextPath := r.Group("/inout-webhook")
	v1 := contextPath.Group("/v1")

	// httoInOut := httpinout.NewHTTP(&httpinout.HTTPConfig{
	// 	Addr:       "localhost:8080/inout-management",
	// 	ApiVersion: "v1",
	// 	Endpoints: &httpinout.Endpoints{
	// 		AddExpenseEndpoint: "/add-expense",
	// 		AddIncomeEndpoint:  "/add-income",
	// 	},
	// })
	bot, err := linebot.New("", "")
	if err != nil {
		panic(err)
	}
	lineService := coreline.NewService()
	lineHandler := linehandleer.NewLineHandler(bot, lineService.EventMessage)
	v1.POST("/line", lineHandler.LineHandler)

	return r
}
