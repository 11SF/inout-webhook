package linehandleer

import (
	"net/http"

	"github.com/11SF/go-common/response"
	coreline "github.com/11SF/inout-webhook/pkg/v1/core/line_event"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineHandler struct {
	linebot      *linebot.Client
	eventMessage coreline.EventMessageFunc
}

func NewLineHandler(linebot *linebot.Client, eventMessage coreline.EventMessageFunc) *LineHandler {
	return &LineHandler{
		linebot:      linebot,
		eventMessage: eventMessage,
	}
}

func (h *LineHandler) LineHandler(c *gin.Context) {

	events, err := h.linebot.ParseRequest(c.Request)
	if err != nil {
		response.NewGinResponseError(c, http.StatusBadRequest, err)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			h.eventMessage(event)
		}
	}

}
