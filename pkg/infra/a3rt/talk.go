package a3rt

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type smallTalkResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Results []*SmallTalkResult `json:"results"`
}

// SmallTalkResult はTalk APIのレスポンスを表現する構造体です
type SmallTalkResult struct {
	Perplexity float64 `json:"perplexity"`
	Reply      string  `json:"reply"`
}

// SmallTalk はTalk APIを通して会話し、返答を受け取るための振る舞いを表現します
func (client *Client) SmallTalk(ctx context.Context, query string) (*SmallTalkResult, error) {
	v := url.Values{
		"query": []string{query},
	}

	var resp smallTalkResponse
	if err := client.do(ctx, http.MethodPost, "talk/v1/smalltalk", v, &resp); err != nil {
		return nil, err
	}

	if resp.Status != 0 {
		return nil, fmt.Errorf("%d: %s", resp.Status, resp.Message)
	}

	return resp.Results[0], nil
}
