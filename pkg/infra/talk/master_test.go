package talk

import (
	"context"
	"testing"

	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/test"
)

func TestNewClient(t *testing.T) {
	c, _ := NewClient(&environment.TalkMaster{
		SpreadsheetID: "1El3atSDe72aNI4jW1gGy_ttwyaZ2r9Om5ZmRoZY9414",
	})
	s, err := c.SelectResponce(context.Background(), "golangのplaygroud")
	if msg := test.Equal(err, nil, "SelectResponce"); msg != "" {
		t.Error(msg)
	}
	if msg := test.Equal(s, nil, "SelectResponce"); msg != "" {
		t.Error(msg)
	}
}
