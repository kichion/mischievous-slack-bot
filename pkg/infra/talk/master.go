package talk

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/file/s3"

	"golang.org/x/oauth2/google"
	"golang.org/x/xerrors"
	"google.golang.org/api/sheets/v4"
)

type config struct {
	spreadsheetID string
}

// Client はTalkMasterのAPIを取り回すためのクライアントを表現する構造体です
type Client struct {
	*http.Client
	config *config
}

// NewClient はTalkMasterのAPIを取り回すためのクライアントを生成して返します
func NewClient(v *environment.TalkMaster) (*Client, error) {
	c, err := httpClient(v)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: c,
		config: &config{spreadsheetID: v.SpreadsheetID},
	}, nil
}

// SelectResponce は指定されたメッセージに対する返答があれば返します
func (client *Client) SelectResponce(ctx context.Context, msg string) (string, error) {
	resp, err := client.getRange(ctx)
	if err != nil {
		return "", err
	}

	for _, row := range resp.Values {
		target := row[0].(string)
		match := row[1].(string)
		responce := row[2].(string)

		switch match {
		case "equal":
			if target == msg {
				return responce, nil
			}
		case "contains":
			if strings.Contains(msg, target) {
				return responce, nil
			}
		}
	}

	return "", nil
}

func httpClient(v *environment.TalkMaster) (*http.Client, error) {
	s, err := s3.NewClient(v)
	if err != nil {
		log.Println(err)
		return nil, xerrors.Errorf("talk master httpClient error: %v", err)
	}
	data, err := s.GetSecret()
	if err != nil {
		log.Println(err)
		return nil, xerrors.Errorf("talk master httpClient error: %v", err)
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Println(err)
		return nil, xerrors.Errorf("talk master httpClient error: %v", err)
	}

	return conf.Client(context.TODO()), nil
}

func (client *Client) getRange(ctx context.Context) (*sheets.ValueRange, error) {
	sheetService, err := sheets.New(client.Client) // nolint:staticcheck
	if err != nil {
		log.Printf("Unable to retrieve Sheets Client %v", err)
		return nil, xerrors.Errorf("talk master getRange error: %v", err)
	}

	r, err := sheetService.Spreadsheets.Values.Get(client.config.spreadsheetID, "A2:C").Context(ctx).Do()
	if err != nil {
		log.Printf("Unable to get Spreadsheets. %v", err)
		return nil, xerrors.Errorf("talk master getRange error: %v", err)
	}

	return r, nil
}
