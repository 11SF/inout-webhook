package coreline

import (
	"strconv"
	"strings"

	"github.com/11SF/inout-webhook/pkg/v1/datamodel"
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

		plaintText := textMessage.Text
		plaintText = strings.TrimSpace(plaintText)

		if len(plaintText) <= 0 {
			return nil
		}

		data := strings.Split(plaintText, "฿")
		if len(data) != 2 {
			slog.Info("invalid in/ex format")
			return nil
		}
		transactionFlag := data[0][0]
		if string(transactionFlag) == string(s.config.AppConstants.ExpenseFlag) {
			slog.Info("start case add expense")
			amountStr := data[0][1:]
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				slog.Info("fail to parse float")
				return nil
			}

			trans := &datamodel.Transaction{
				Amount:   amount,
				Message:  data[1],
				UserUUID: event.Source.UserID,
			}
			err = s.httpInOut.AddExpenseRequest(trans)
			if err != nil {
				slog.Info("fail to add expense", "with", err.Error())
				return nil
			}
			slog.Info("case add expense success")
			return nil
		}
		if string(transactionFlag) == string(s.config.AppConstants.IncomeFlag) {
			slog.Info("start case add income")
			amountStr := data[0][1:]
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				slog.Info("fail to parse float")
				return nil
			}

			trans := &datamodel.Transaction{
				Amount:   amount,
				Message:  data[1],
				UserUUID: event.Source.UserID,
			}
			err = s.httpInOut.AddIncomeRequest(trans)
			if err != nil {
				slog.Info("fail to add income", "with", err.Error())
				return nil
			}
			return nil
		}
	}

	slog.Info("case add income success")
	return nil
}
