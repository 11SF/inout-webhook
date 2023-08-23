package httpinout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/11SF/go-common/response"
	"github.com/11SF/inout-webhook/pkg/v1/datamodel"
	"golang.org/x/exp/slog"
)

type AddExpenseRequestFunc func(trans *datamodel.Transaction) error

func (h *httpInOut) AddExpenseRequest(trans *datamodel.Transaction) error {

	slog.Info("start http request to add expense")

	transByte, err := json.Marshal(trans)
	if err != nil {
		slog.Info("fail to marshal transaction", "with", err.Error())
		return response.NewError("IN500", err.Error())
	}

	request, err := http.NewRequest("POST", h.getUrl(h.config.Endpoints.AddExpenseEndpoint), bytes.NewBuffer(transByte))
	if err != nil {
		slog.Info("fail to marshal transaction", "with", err.Error())
		return response.NewError("TP500", err.Error())
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		slog.Info("fail to marshal transaction", "with", err.Error())
		return response.NewError("TP500", err.Error())
	}
	if res.StatusCode >= 300 {
		slog.Info("server fail to add expense", "with error code", res.StatusCode)
		return response.NewError("TP500", fmt.Sprintf("server fail to add expense with error code %v", res.StatusCode))
	}

	return nil
}
