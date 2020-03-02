package mention

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kichion/mischievous-slack-bot/pkg/domain/valueobject/responce"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/a3rt"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

// Talk は会話的な返事行うための振る舞いです
func Talk(e *slackevents.AppMentionEvent, v *environment.Variable) (events.APIGatewayProxyResponse, error) {
	api := slack.New(v.Slack.OAuthAccessToken)
	talkAPI := a3rt.NewTalkClient(&v.A3RT)

	res, err := talkAPI.SmallTalk(context.Background(), strings.Replace(e.Text, v.Slack.BotMention+" ", "", 1)) // メンション部分を除去する
	if err != nil {
		log.Print(err)
		return responce.NewGateway(http.StatusBadRequest), err
	}

	_, _, err = api.PostMessage(
		e.Channel,
		slack.MsgOptionText(
			res.Reply,
			false,
		),
	)
	if err != nil {
		log.Print(err)
		return responce.NewGateway(http.StatusBadRequest), err
	}

	return responce.NewGateway(http.StatusOK), nil
}
