package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kichion/mischievous-slack-bot/pkg/actions"
	"github.com/kichion/mischievous-slack-bot/pkg/actions/mention"
	"github.com/kichion/mischievous-slack-bot/pkg/domain/valueobject/responce"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/kichion/mischievous-slack-bot/pkg/service/parser"
	"github.com/kichion/mischievous-slack-bot/pkg/service/verifer"
	"github.com/slack-go/slack/slackevents"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	v, err := environment.New()
	if err != nil {
		log.Print(err.Error())
		return responce.NewGateway(http.StatusInternalServerError), err
	}

	event, err := parser.LambdaEventToSlackEvent(request, v)
	if err != nil {
		log.Print(err.Error())
		return responce.NewGateway(http.StatusInternalServerError), err
	}

	if actions.IsURLVerification(event) {
		return actions.URLVerification(request.Body)
	}
	if actions.IsUndefined(event) {
		return responce.NewGateway(http.StatusBadRequest), nil
	}

	if err := verifer.NewSecrets(request, v); err != nil {
		log.Print(err.Error())
		return responce.NewGateway(http.StatusBadRequest), err
	}

	switch e := event.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		if e.Text == v.Slack.BotMention+" ping" {
			return mention.Monotonous(e, v)
		}
		if strings.Contains(e.Text, " typo ") {
			return mention.Proofreading(e, v)
		}
		if e.Text == v.Slack.BotMention+" curry order" {
			return mention.CurryOrder(e, v)
		}
		return mention.Talk(e, v)
	}

	return responce.NewGateway(http.StatusBadRequest), nil
}

func main() {
	lambda.Start(handler)
}
