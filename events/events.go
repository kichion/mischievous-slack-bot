package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var api = slack.New(os.Getenv("BOT_OAUTH_ACCESS_TOKEN"))

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	reqBody := request.Body
	eventsAPIEvent, err := slackevents.ParseEvent(
		json.RawMessage(reqBody),
		slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: os.Getenv("VERIFICATION_TOKEN")}),
	)
	if err != nil {
		log.Print(err)
		return events.APIGatewayProxyResponse{}, err
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(reqBody), &r)
		if err != nil {
			log.Print(err)
			return events.APIGatewayProxyResponse{}, err
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       r.Challenge,
		}, nil
	}
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		log.Print(innerEvent.Type)
		var err error
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			_, _, err = api.PostMessage(ev.Channel, slack.MsgOptionText(fmt.Sprintf("Yes, hello."), false)) // nolint:errcheck
		}
		if err != nil {
			log.Print(err)
			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body:       "Bad Request",
			}, nil
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Success",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
