package anyapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"net/http"
)

type AnyApi api.Service

func (a *AnyApi) AnyApi(ctx context.Context, method string, bizContent any) (*http.Response, *core.Response, error) {

	reqBody, err := a.BuildAnyApiParams(bizContent)
	if err != nil {
		return nil, nil, err
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, nil, err
	}

	url := fmt.Sprintf("%s%s", a.Config.BaseURL, method)

	httpResponse, bodyBytes, err := a.Request(ctx, http.MethodPost, url, reqBodyBytes)
	if err != nil {
		return nil, nil, err
	}

	res, err := core.BuildResponse(bodyBytes)
	if err != nil {
		return httpResponse, nil, err
	}

	return httpResponse, res, nil
}
