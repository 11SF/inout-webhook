package coreline

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/exp/slog"
)

type EventMessageFunc func(event *linebot.Event) error

func (s *service) EventMessage(event *linebot.Event) error {

	slog.Info("start event message")
	if event.Type == linebot.EventTypeMessage {
		textMessage, ok := event.Message.(*linebot.TextMessage)
		if !ok {
			return nil
		}
		slog.Info(event.Source.UserID)
		slog.Info(textMessage.Text)
	}

	return nil
}
