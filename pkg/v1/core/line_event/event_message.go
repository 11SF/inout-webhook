package coreline

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/11SF/inout-webhook/pkg/v1/datamodel"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/exp/slog"
)

type EventMessageFunc func(event *linebot.Event) error

func (s *service) EventMessage(event *linebot.Event) error {

	slog.Info("start event message")
	textMessage, ok := event.Message.(*linebot.TextMessage)
	if !ok {
		return nil
	}
	slog.Info(textMessage.Text)

	plaintText := textMessage.Text
	plaintText = strings.TrimSpace(plaintText)

	var transactionFlag string
	var trans *datamodel.Transaction

	transactionWithMessage, _ := regexp.MatchString(`([+-])(\d+(\.\d{2})?)฿(.+)`, plaintText)
	if transactionWithMessage {
		fmt.Println("transactionWithMessage")
		transactionFlag = string(plaintText[0])
		data := strings.Split(plaintText, "฿")
		amountStr := data[0][1:]
		amount, _ := strconv.ParseFloat(amountStr, 64)

		trans = &datamodel.Transaction{
			Amount:   amount,
			Message:  strings.TrimSpace(data[1]),
			UserUUID: event.Source.UserID,
		}
	}
	transactionWithOutMessage, _ := regexp.MatchString(`([+-])(\d+(\.\d{2})?)฿$`, plaintText)
	if transactionWithOutMessage {
		fmt.Println("transactionWithOutMessage")
		transactionFlag = string(plaintText[0])
		amountStr := strings.TrimRight(plaintText[1:], "฿")
		amountStr = strings.TrimSpace(amountStr)
		amount, _ := strconv.ParseFloat(amountStr, 64)

		trans = &datamodel.Transaction{
			Amount:   amount,
			UserUUID: event.Source.UserID,
		}
	}

	if transactionFlag == string(s.config.AppConstants.ExpenseFlag) {
		slog.Info("start case add expense")
		err := s.httpInOut.AddExpenseRequest(trans)
		if err != nil {
			slog.Info("fail to add expense", "with", err.Error())
			return nil
		}
		slog.Info("case add expense success")
		return nil

	}
	if transactionFlag == string(s.config.AppConstants.IncomeFlag) {
		slog.Info("start case add income")
		err := s.httpInOut.AddIncomeRequest(trans)
		if err != nil {
			slog.Info("fail to add income", "with", err.Error())
			return nil
		}
		slog.Info("case add income success")
		return nil

	}

	slog.Info("unexpected case")
	return nil
}
