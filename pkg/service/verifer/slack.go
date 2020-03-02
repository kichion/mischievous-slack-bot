package verifer

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/slack-go/slack"

	"golang.org/x/xerrors"
)

// NewSecrets はSiging Secretとリクエストのbodyやtimestampを組み合わせて生成されたものを検証します
func NewSecrets(request events.APIGatewayProxyRequest, v *environment.Variable) error {
	headers := convertHeaders(request.Headers)
	sv, err := slack.NewSecretsVerifier(headers, v.Slack.SigningSecret)
	if err != nil {
		return xerrors.Errorf("NewSecrets error: %v", err)
	}
	sv.Write([]byte(request.Body)) // nolint
	return sv.Ensure()
}

func convertHeaders(headers map[string]string) http.Header {
	h := http.Header{}
	for key, value := range headers {
		h.Set(key, value)
	}
	return h
}
