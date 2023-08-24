package httpinout

import "fmt"

type HTTPInOut struct {
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

func NewHTTP(config *HTTPConfig) *HTTPInOut {
	return &HTTPInOut{
		config: config,
	}
}

func (h *HTTPInOut) getUrl(endpoint string) string {
	return fmt.Sprintf("%v/%v%v", h.config.Addr, h.config.ApiVersion, endpoint)
}
