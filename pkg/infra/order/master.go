package order

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/file/s3"

	"golang.org/x/oauth2/google"
	"golang.org/x/xerrors"
	"google.golang.org/api/sheets/v4"
)

type config struct {
	spreadsheetID string
}

// Client はOrderMasterのAPIを取り回すためのクライアントを表現する構造体です
type Client struct {
	*http.Client
	config *config
}

// Item は注文の情報をまとめた構造体です
type Item struct {
	Price int
	Icon  string
	Name  string
	Desc  string
	UUID  string
}

// NewClient はOrderMasterのAPIを取り回すためのクライアントを生成して返します
func NewClient(v *environment.OrderMaster) (*Client, error) {
	c, err := httpClient(v)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: c,
		config: &config{spreadsheetID: v.SpreadsheetID},
	}, nil
}

// OrderList は指定されたメッセージに対する返答があれば返します
func (client *Client) OrderList(ctx context.Context) ([]*Item, error) {
	resp, err := client.getRange(ctx)
	if err != nil {
		return []*Item{}, err
	}

	orders := []*Item{}
	for _, row := range resp.Values {
		p := row[0].(string)
		price, _ := strconv.Atoi(p)
		orders = append(orders, &Item{
			Price: price,
			Icon:  row[1].(string),
			Name:  row[2].(string),
			Desc:  row[3].(string),
			UUID:  row[4].(string),
		})
	}

	return orders, nil
}

func httpClient(v *environment.OrderMaster) (*http.Client, error) {
	s, err := s3.NewClient(v)
	if err != nil {
		log.Println(err)
		return nil, xerrors.Errorf("order master httpClient error: %v", err)
	}
	data, err := s.GetSecret()
	if err != nil {
		log.Println(err)
		return nil, xerrors.Errorf("order master httpClient error: %v", err)
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Println(err)
		return nil, xerrors.Errorf("order master httpClient error: %v", err)
	}

	return conf.Client(context.TODO()), nil
}

func (client *Client) getRange(ctx context.Context) (*sheets.ValueRange, error) {
	sheetService, err := sheets.New(client.Client) // nolint:staticcheck
	if err != nil {
		log.Printf("Unable to retrieve Sheets Client %v", err)
		return nil, xerrors.Errorf("order master getRange error: %v", err)
	}

	r, err := sheetService.Spreadsheets.Values.Get(client.config.spreadsheetID, "order!A2:E").Context(ctx).Do()
	if err != nil {
		log.Printf("Unable to get Spreadsheets. %v", err)
		return nil, xerrors.Errorf("order master getRange error: %v", err)
	}

	return r, nil
}
