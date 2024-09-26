package key

import (
	"encoding/json"
	"fmt"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
)

type Key api.Service

func (a *Key) Order(request *OrderRequest) (*Response, error) {
	b, err := a.Request(orderMethod, request)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *Key) Query(request *QueryRequest) (*Response, error) {
	b, err := a.Request(queryMethod, request)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *Key) Discard(request *DiscardRequest) (*Response, error) {
	b, err := a.Request(discardMethod, request)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *Key) Notify(notify *Notify) (*Reply, error) {
	b := a.Verifier.Verify(notify.SignStr(), notify.Sign)
	if !b {
		return nil, fmt.Errorf("verify sign fail")
	}
	return notify.Data, nil
}
