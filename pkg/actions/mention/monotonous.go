package mention

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kichion/mischievous-slack-bot/pkg/domain/valueobject/responce"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

// Monotonous は単調な返事行うための振る舞いです
func Monotonous(e *slackevents.AppMentionEvent, v *environment.Variable) (events.APIGatewayProxyResponse, error) {
	api := slack.New(v.Slack.OAuthAccessToken)

	_, _, err := api.PostMessage(
		e.Channel,
		slack.MsgOptionText(
			fmt.Sprintf("hi, hello."),
			false,
		),
	)
	if err != nil {
		log.Print(err)
		return responce.NewGateway(http.StatusBadRequest), err
	}

	return responce.NewGateway(http.StatusOK), nil
}
