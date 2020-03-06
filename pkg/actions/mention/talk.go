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
	"github.com/kichion/mischievous-slack-bot/pkg/infra/talk"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"golang.org/x/xerrors"
)

// Talk は会話的な返事行うための振る舞いです
func Talk(e *slackevents.AppMentionEvent, v *environment.Variable) (events.APIGatewayProxyResponse, error) {
	msg, err := selectMessage(e, v)
	if err != nil {
		printErr := xerrors.Errorf("Talk error: %v", err)
		log.Print(printErr)
		return responce.NewGateway(http.StatusBadRequest), printErr
	}

	api := slack.New(v.Slack.OAuthAccessToken)

	if _, _, err = api.PostMessage(
		e.Channel,
		slack.MsgOptionText(
			msg,
			false,
		),
	); err != nil {
		printErr := xerrors.Errorf("Talk error: %v", err)
		log.Print(printErr)
		return responce.NewGateway(http.StatusBadRequest), printErr
	}

	return responce.NewGateway(http.StatusOK), nil
}

func selectMessage(e *slackevents.AppMentionEvent, v *environment.Variable) (string, error) {
	if message, err := definitiveResponce(e, v); message != "" && err == nil {
		return message, nil
	}

	return flexibleResponce(e, v)
}

func definitiveResponce(e *slackevents.AppMentionEvent, v *environment.Variable) (string, error) {
	api, err := talk.NewClient(&v.TalkMaster)
	if err != nil {
		return "", err
	}

	msg, err := api.SelectResponce(
		context.Background(),
		getTalkRequestMessage(e.Text, v.Slack.BotMention),
	)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func flexibleResponce(e *slackevents.AppMentionEvent, v *environment.Variable) (string, error) {
	api := a3rt.NewTalkClient(&v.A3RT)

	res, err := api.SmallTalk(
		context.Background(),
		getTalkRequestMessage(e.Text, v.Slack.BotMention),
	)
	if err != nil {
		return "", err
	}
	return res.Reply, nil
}

func getTalkRequestMessage(text string, mention string) string {
	if strings.Contains(text, mention) {
		return strings.Replace(text, mention+" ", "", 1) // メンション部分を除去する
	}
	return text
}
