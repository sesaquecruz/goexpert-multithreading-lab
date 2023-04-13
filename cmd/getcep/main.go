package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/sesaquecruz/goexpert-multithreading-lab/internal/service"
)

const (
	Timeout = 1 * time.Second

	ApiName1 = "apicep"
	ApiUrl1  = "https://cdn.apicep.com/file/apicep/%s.json"

	ApiName2 = "viacep"
	ApiUrl2  = "http://viacep.com.br/ws/%s/json/"
)

var CepPattern = regexp.MustCompile(`^[0-9]{5}-[0-9]{3}$`)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("fatal error: pass a cep as argument (ex: getcep xxxxx-xxx)")
		return
	}

	cep := os.Args[1]
	if !CepPattern.MatchString(cep) {
		fmt.Println("fatal error: invalid cep format (ex: xxxxx-xxx)")
		return
	}

	service1 := service.NewCepService(ApiName1, ApiUrl1)
	service2 := service.NewCepService(ApiName2, ApiUrl2)

	response1 := make(chan service.Response)
	response2 := make(chan service.Response)

	ctx := context.Background()

	go GetCep(ctx, cep, service1, response1)
	go GetCep(ctx, cep, service2, response2)

	select {
	case res := <-response1:
		fmt.Printf("result: %s\n", res)
	case res := <-response2:
		fmt.Printf("result: %s\n", res)
	case <-time.After(Timeout):
		fmt.Println("result: timeout reached")
	}
}

func GetCep(ctx context.Context, cep string, service *service.CepService, response chan<- service.Response) {
	res, err := service.Get(ctx, cep)
	if err != nil {
		fmt.Printf("log error: %s\n", err)
	} else {
		response <- *res
		close(response)
	}
}
