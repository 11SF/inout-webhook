package httpinout

import "fmt"

type httpInOut struct {
	config *HTTPConfig
}

type HTTPConfig struct {
	Addr       string
	ApiVersion string
	Endpoints  *Endpoints
}

type Endpoints struct {
	AddIncomeEndpoint  string
	AddExpenseEndpoint string
}

func NewHTTP(config *HTTPConfig) *httpInOut {
	return &httpInOut{
		config: config,
	}
}

func (h *httpInOut) getUrl(endpoint string) string {
	return fmt.Sprintf("%v/%v%v", h.config.Addr, h.config.ApiVersion, endpoint)
}
