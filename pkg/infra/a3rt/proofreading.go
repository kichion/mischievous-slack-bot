package a3rt

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ProofreadingResponse はタイポ検査の結果を表現する構造体です
type ProofreadingResponse struct {
	ResultID           string `json:"resultID"`
	Status             int    `json:"status"`
	Message            string `json:"message"`
	InputSentence      string `json:"inputSentence"`
	NormalizedSentence string `json:"normalizedSentence"`
	CheckedSentence    string `json:"checkedSentence"`

	Alerts []*ProofreadingAlerts `json:"alerts"`
}

// ProofreadingAlerts は指定内容を表現する構造体です
type ProofreadingAlerts struct {
	Pos         int      `json:"pos"`
	Word        string   `json:"word"`
	Score       float64  `json:"score"`
	Suggestions []string `json:"suggestions"`
}

// Proofreading はProofreading APIを通して文章を検査し、返答を受け取るための振る舞いを表現します
func (client *Client) Proofreading(ctx context.Context, sentence string) (*ProofreadingResponse, error) {
	v := url.Values{
		"sentence": []string{sentence},
	}

	var resp ProofreadingResponse
	if err := client.do(ctx, http.MethodPost, "proofreading/v2/typo", v, &resp); err != nil {
		return nil, err
	}

	if resp.Status != 0 && resp.Status != 1 {
		return nil, fmt.Errorf("%d: %s", resp.Status, resp.Message)
	}

	return &resp, nil
}
