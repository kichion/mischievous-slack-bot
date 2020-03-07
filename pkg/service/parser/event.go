package parser

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

// LambdaEventToSlackEvent はLambda EventをSlack Eventへ変換します
func LambdaEventToSlackEvent(request events.APIGatewayProxyRequest, v *environment.Variable) (slackevents.EventsAPIEvent, error) {
	return slackevents.ParseEvent(
		json.RawMessage(request.Body),
		slackevents.OptionVerifyToken(
			&slackevents.TokenComparator{
				VerificationToken: v.Slack.VerificationToken,
			},
		),
	)
}

// LambdaEventToInteractionCallback はLambda EventをInteraction Callbackへ変換します
func LambdaEventToInteractionCallback(request events.APIGatewayProxyRequest, v *environment.Variable) (*slack.InteractionCallback, error) {
	str, _ := url.QueryUnescape(request.Body)
	str = strings.Replace(str, "payload=", "", 1)
	var payload slack.InteractionCallback
	if err := json.Unmarshal([]byte(str), &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}
