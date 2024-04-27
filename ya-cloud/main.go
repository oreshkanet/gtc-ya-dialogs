package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oreshkanet/gtc-ya-dialogs/alice"
	"github.com/oreshkanet/gtc-ya-dialogs/logger"
	"time"
)

func Handler(ctx context.Context, req *alice.Request) (*alice.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	aliceApp, err := getAliceApp()
	if err != nil {
		return nil, err
	}
	return aliceApp.Handle(ctx, req)
}

// ?client_id=test&redirect_uri=https%3A%2F%2Fsocial.yandex.net%2Fbroker%2Fredirect&response_type=code&state=https%3A%2F%2Fsocial.yandex.ru%2Fbroker2%2Fauthz_in_web%2F1a02abf8b7ed449c9a4a69cd07000560%2Fcallback: {}

// Структура запроса API Gateway v1
type APIGatewayRequest struct {
	OperationID string `json:"operationId"`
	Resource    string `json:"resource"`

	HTTPMethod string `json:"httpMethod"`

	Path           string            `json:"path"`
	PathParameters map[string]string `json:"pathParameters"`

	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`

	QueryStringParameters           map[string]string   `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`

	Parameters           map[string]string   `json:"parameters"`
	MultiValueParameters map[string][]string `json:"multiValueParameters"`

	Body            []byte `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded,omitempty"`

	RequestContext interface{} `json:"requestContext"`
}

// Структура ответа API Gateway v1
type APIGatewayResponse struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded,omitempty"`
}

type Request1 struct {
	Name string `json:"name"`
}

func Greet(ctx context.Context, event *APIGatewayRequest) (*APIGatewayResponse, error) {
	ctx, err := initLogging(ctx)
	if err != nil {
		//return nil, err
	}

	//defer log.Sync(ctx)

	logger.Info(ctx, "Handler")

	// В журнале будет напечатано название HTTP-метода, с помощью которого осуществлен запрос, а также путь
	v, err := json.Marshal(event.PathParameters)
	logger.Info(ctx, string(v))
	v, err = json.Marshal(event.QueryStringParameters)
	logger.Info(ctx, string(v))
	logger.Info(ctx, event.Resource)

	req := &Request{}

	// Поле event.Body запроса преобразуется в объект типа Request для получения переданного имени
	if err := json.Unmarshal(event.Body, &req); err != nil {
		return nil, fmt.Errorf("an error has occurred when parsing body: %v", err)
	}

	// Тело ответа.
	return &APIGatewayResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Hello, %s", req.Meta.ClientId),
	}, nil
}
