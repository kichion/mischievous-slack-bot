package parser

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
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
