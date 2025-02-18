package v2

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
	"net/http"
)

const (
	orderMethod = "/openapi/v2/key/order"
	queryMethod = "/openapi/v2/key/query"
)

type Key api.Service

func (k *Key) Order(ctx context.Context, request *OrderRequest) (*http.Response, *Response, error) {
	httpResponse, bodyBytes, err := k.Post(ctx, orderMethod, request)
	if err != nil {
		return nil, nil, err
	}

	res, err := response(bodyBytes)
	if err != nil {
		return httpResponse, nil, err
	}

	return httpResponse, res, nil
}

func (k *Key) Query(ctx context.Context, request *QueryRequest) (*http.Response, *Response, error) {
	httpResponse, bodyBytes, err := k.Post(ctx, queryMethod, request)
	if err != nil {
		return nil, nil, err
	}

	res, err := response(bodyBytes)
	if err != nil {
		return httpResponse, nil, err
	}

	return httpResponse, res, nil
}
