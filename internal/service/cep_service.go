package service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Response struct {
	ApiName string
	Data    string
}

type CepService struct {
	ApiName string
	ApiUrl  string
}

func NewCepService(apiName string, apiUrl string) *CepService {
	return &CepService{
		ApiName: apiName,
		ApiUrl:  apiUrl,
	}
}

func (a *CepService) Get(ctx context.Context, cep string) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(a.ApiUrl, cep), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s - %d", a.ApiName, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Response{a.ApiName, strings.ReplaceAll(string(body), "\n", "")}, nil
}
