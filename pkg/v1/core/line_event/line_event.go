package coreline

import (
	"github.com/11SF/inout-webhook/configs"
	httpinout "github.com/11SF/inout-webhook/pkg/v1/core/http"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type service struct {
	httpInOut *httpinout.HTTPInOut
	config    *configs.Config
}

type Service interface {
	EventMessage(event *linebot.Event) error
}

func NewService(httpInOut *httpinout.HTTPInOut, config *configs.Config) *service {
	return &service{httpInOut, config}
}
