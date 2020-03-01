package a3rt

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
)

type config struct {
	baseURL string
	apiKey  string
}

// Client はA3RTのAPIを取り回すためのクライアントを表現する構造体です
type Client struct {
	*http.Client
	config *config
}

const baseURL = "https://api.a3rt.recruit-tech.co.jp"

// NewClient はA3RTのAPIを取り回すためのクライアントを生成して返します
func NewClient(v *environment.A3RT) *Client {
	return &Client{
		Client: http.DefaultClient,
		config: &config{
			baseURL: baseURL,
			apiKey:  v.TalkAPIKey,
		},
	}
}

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
	v := url.Values{}
	v.Add("query", query)

	var resp smallTalkResponse
	if err := client.do(ctx, http.MethodPost, "talk/v1/smalltalk", v, &resp); err != nil {
		return nil, err
	}

	if resp.Status != 0 {
		return nil, fmt.Errorf("%d: %s", resp.Status, resp.Message)
	}

	return resp.Results[0], nil
}

func (client *Client) do(ctx context.Context, method string, uri string, params url.Values, res interface{}) error {
	u, err := url.Parse(client.config.baseURL)
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, uri)

	params.Add("apikey", client.config.apiKey)

	var body io.Reader
	switch method {
	case http.MethodGet:
		u.RawQuery = params.Encode()
	default:
		body = strings.NewReader(params.Encode())
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return err
	}
	_ = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if res == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(&res)
}
