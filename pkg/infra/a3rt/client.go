package a3rt

import (
	"context"
	"encoding/json"
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

// NewTalkClient はA3RTのTalk APIを取り回すためのクライアントを生成して返します
func NewTalkClient(v *environment.A3RT) *Client {
	return &Client{
		Client: http.DefaultClient,
		config: &config{
			baseURL: baseURL,
			apiKey:  v.TalkAPIKey,
		},
	}
}

// NewProofreadingClient はA3RTのProofreading APIを取り回すためのクライアントを生成して返します
func NewProofreadingClient(v *environment.A3RT) *Client {
	return &Client{
		Client: http.DefaultClient,
		config: &config{
			baseURL: baseURL,
			apiKey:  v.ProofreadingAPIKey,
		},
	}
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
