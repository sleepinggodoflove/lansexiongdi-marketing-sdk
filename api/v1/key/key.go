package key

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"io"
	"net/http"
)

const (
	orderMethod   = "/openapi/v1/key/order"
	queryMethod   = "/openapi/v1/key/query"
	discardMethod = "/openapi/v1/key/discard"
)

type Key api.Service

func (k *Key) Order(ctx context.Context, request *OrderRequest) (*http.Response, *core.Response, error) {

	httpResponse, bodyBytes, err := k.Post(ctx, orderMethod, request)
	if err != nil {
		return nil, nil, err
	}

	res, err := core.BuildResponse(bodyBytes)
	if err != nil {
		return httpResponse, nil, err
	}

	return httpResponse, res, nil
}

func (k *Key) Query(ctx context.Context, request *QueryRequest) (*http.Response, *core.Response, error) {

	httpResponse, bodyBytes, err := k.Post(ctx, queryMethod, request)
	if err != nil {
		return nil, nil, err
	}

	res, err := core.BuildResponse(bodyBytes)
	if err != nil {
		return httpResponse, nil, err
	}

	return httpResponse, res, nil
}

func (k *Key) Discard(ctx context.Context, request *DiscardRequest) (*http.Response, *core.Response, error) {

	httpResponse, bodyBytes, err := k.Post(ctx, discardMethod, request)
	if err != nil {
		return nil, nil, err
	}

	res, err := core.BuildResponse(bodyBytes)
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

func (k *Key) CallBack(ctx context.Context, req *http.Request) (*NotifyData, error) {

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var n *Notify
	if err = json.Unmarshal(body, &n); err != nil {
		return nil, err
	}

	sign := req.Header.Get("Sign")
	timestamp := req.Header.Get("Timestamp")

	if sign == "" {
		return k.Notify(ctx, n)
	}

	if timestamp == "" {
		return nil, fmt.Errorf("timestamp is empty")
	}

	ciphertext, err := k.GetCiphertext(&n.Data)
	if err != nil {
		return nil, err
	}
	fmt.Printf(ciphertext)

	if !k.Verify(timestamp, ciphertext, sign) {
		return nil, fmt.Errorf("call back verify sign fail")
	}

	return &n.Data, nil
}
