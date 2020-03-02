package mention

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kichion/mischievous-slack-bot/pkg/domain/valueobject/responce"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/a3rt"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"

	"golang.org/x/xerrors"
)

// Proofreading はタイポチェック行うための振る舞いです
func Proofreading(e *slackevents.AppMentionEvent, v *environment.Variable) (events.APIGatewayProxyResponse, error) {
	slackAPI := slack.New(v.Slack.OAuthAccessToken)
	proofreadingAPI := a3rt.NewProofreadingClient(&v.A3RT)

	res, err := proofreadingAPI.Proofreading(
		context.Background(),
		getProofRequestMessage(e.Text, v.Slack.BotMention),
	)
	if err != nil {
		printErr := xerrors.Errorf("Proofreading error: %v", err)
		log.Print(printErr)
		return responce.NewGateway(http.StatusBadRequest), printErr
	}

	_, _, err = slackAPI.PostMessage(
		e.Channel,
		slack.MsgOptionText(
			buildTypoAlert(res),
			false,
		),
	)
	if err != nil {
		printErr := xerrors.Errorf("Proofreading error: %v", err)
		log.Print(printErr)
		return responce.NewGateway(http.StatusBadRequest), printErr
	}

	return responce.NewGateway(http.StatusOK), nil
}

func buildTypoAlert(res *a3rt.ProofreadingResponse) string {
	if len(res.Alerts) == 0 {
		return "すばらしい！問題ないと思います！"
	}

	msg := fmt.Sprintf("誤りがありそうです…\n```%s```\n", res.CheckedSentence)
	for _, alert := range res.Alerts {
		msg += fmt.Sprintf("`%s` は 「%s」のいずれかが正しいと思われます (精度: %s)\n", alert.Word, strings.Join(alert.Suggestions, ","), fmt.Sprintf("%.2f", alert.Score))
	}
	return msg
}

func getProofRequestMessage(text string, mention string) string {
	if strings.Contains(text, mention) {
		return strings.Replace(text, mention+" typo ", "", 1) // メンション部分を除去する
	}
	return strings.Replace(text, "typo ", "", 1)
}
