package actions

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kichion/mischievous-slack-bot/pkg/domain/valueobject/responce"
	"github.com/slack-go/slack/slackevents"
)

// URLVerification はSlack Bot AppのURL検証のためのアクションを行います
func URLVerification(body string) (events.APIGatewayProxyResponse, error) {
	var r *slackevents.ChallengeResponse
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		log.Print(err.Error())
		return responce.NewGateway(http.StatusBadRequest), err
	}

	return responce.NewGatewayChallenge(r.Challenge), nil
}
