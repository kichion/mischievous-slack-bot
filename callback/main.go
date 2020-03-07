package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kichion/mischievous-slack-bot/pkg/actions/callback"
	"github.com/kichion/mischievous-slack-bot/pkg/domain/valueobject/responce"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/kichion/mischievous-slack-bot/pkg/service/parser"
	"github.com/kichion/mischievous-slack-bot/pkg/service/verifer"
	"github.com/slack-go/slack"
	"golang.org/x/xerrors"
)

func callBackHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	v, err := environment.New()
	if err != nil {
		log.Print(err.Error())
		return responce.NewGateway(http.StatusInternalServerError), err
	}

	if sErr := verifer.NewSecrets(request, v); sErr != nil {
		log.Print(sErr.Error())
		return responce.NewGateway(http.StatusBadRequest), sErr
	}

	log.Println(request.Body)
	payload, err := parser.LambdaEventToInteractionCallback(request, v)
	if err != nil {
		log.Print(err.Error())
		return responce.NewGateway(http.StatusInternalServerError), err
	}

	// TODO: switch Block Kit actions
	blocks := callback.CurrySelect(payload)

	api := slack.New(v.Slack.OAuthAccessToken)

	if _, _, _, err = api.UpdateMessage(
		payload.Channel.ID,
		payload.Message.Timestamp,
		slack.MsgOptionBlocks(blocks...),
	); err != nil {
		printErr := xerrors.Errorf("curry callback error: %v", err)
		log.Print(printErr)
		return responce.NewGateway(http.StatusBadRequest), printErr
	}

	return responce.NewGateway(http.StatusOK), nil
}

func main() {
	lambda.Start(callBackHandler)
}
