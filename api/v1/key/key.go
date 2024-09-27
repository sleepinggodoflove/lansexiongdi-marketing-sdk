package key

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
)

type Key api.Service

func (a *Key) Order(ctx context.Context, request *OrderRequest) (*Response, error) {
	b, err := a.Request(ctx, orderMethod, request)
	if err != nil {
		return nil, err
	}
	var response Response
	if err = json.Unmarshal(b, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *Key) Query(ctx context.Context, request *QueryRequest) (*Response, error) {
	b, err := a.Request(ctx, queryMethod, request)
	if err != nil {
		return nil, err
	}
	var response Response
	if err = json.Unmarshal(b, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *Key) Discard(ctx context.Context, request *DiscardRequest) (*Response, error) {
	b, err := a.Request(ctx, discardMethod, request)
	if err != nil {
		return nil, err
	}
	var response Response
	if err = json.Unmarshal(b, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *Key) Notify(_ context.Context, notify *Notify) (*Reply, error) {
	if !a.Verifier.Verify(notify.SignStr(), notify.Sign) {
		return nil, fmt.Errorf("verify sign fail")
	}
	return notify.Data, nil
}
