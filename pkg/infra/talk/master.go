package talk

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"golang.org/x/xerrors"
)

type config struct {
	baseURL string
	token   string
}

// Client はTalkMasterのAPIを取り回すためのクライアントを表現する構造体です
type Client struct {
	*http.Client
	config *config
}

// NewClient はTalkMasterのAPIを取り回すためのクライアントを生成して返します
func NewClient(v *environment.TalkMaster) *Client {
	return &Client{
		Client: http.DefaultClient,
		config: &config{
			baseURL: v.BaseURL,
			token:   v.Token,
		},
	}
}

type talkMasterResponce struct {
	Target   string `json:"target"`
	Match    string `json:"match"`
	Responce string `json:"responce"`
}

// SelectResponce は指定されたメッセージに対する返答があれば返します
func (client *Client) SelectResponce(ctx context.Context, msg string) (string, error) {
	var resp []talkMasterResponce
	if err := client.do(ctx, "", url.Values{}, &resp); err != nil {
		return "", err
	}

	if resp == nil {
		return "", errors.New("data undefined")
	}

	for _, mas := range resp {
		switch mas.Match {
		case "equal":
			if mas.Target == msg {
				return mas.Responce, nil
			}
		case "contains":
			if strings.Contains(msg, mas.Target) {
				return mas.Responce, nil
			}
		}
	}

	return "", nil
}

func (client *Client) do(ctx context.Context, uri string, params url.Values, res interface{}) error {
	u, err := url.Parse(client.config.baseURL)
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, uri)

	params.Add("token", client.config.token)

	var body io.Reader
	u.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), body)
	if err != nil {
		return xerrors.Errorf("", err)
	}
	_ = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

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
