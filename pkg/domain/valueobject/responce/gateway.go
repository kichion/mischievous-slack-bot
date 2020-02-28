package responce

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// NewGateway は指定のステータスを返すレスポンスを生成します
func NewGateway(status int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}
}

// NewGatewayChallenge は指定のChallengeを返すレスポンスを生成します
func NewGatewayChallenge(challenge string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       challenge,
	}
}
