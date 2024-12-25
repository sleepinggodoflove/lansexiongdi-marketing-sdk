package key

import (
	"context"
	"fmt"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
	"net/http"
)

const (
	orderMethod   = "/openapi/v1/key/order"
	queryMethod   = "/openapi/v1/key/query"
	discardMethod = "/openapi/v1/key/discard"
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

func (k *Key) Discard(ctx context.Context, request *DiscardRequest) (*http.Response, *Response, error) {
	httpResponse, bodyBytes, err := k.Post(ctx, discardMethod, request)
	if err != nil {
		return nil, nil, err
	}

	res, err := response(bodyBytes)
	if err != nil {
		return httpResponse, nil, err
	}

	return httpResponse, res, nil
}

func (k *Key) Notify(_ context.Context, n *Notify) (*NotifyData, error) {
	if !k.CryptographySuite.Verifier.Verify(n.SignString(), n.Sign) {
		return nil, fmt.Errorf("verify sign fail")
	}
	return &n.Data, nil
}
