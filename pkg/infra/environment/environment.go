package environment

import (
	"golang.org/x/xerrors"

	"github.com/kelseyhightower/envconfig"
)

// Variable 環境変数の情報を保持する構造体です
type Variable struct {
	Slack
}

// Slack Slackにおける情報を保持する構造体です
type Slack struct {
	VerificationToken string `envconfig:"VERIFICATION_TOKEN" required:"true"`
	OAuthAccessToken  string `envconfig:"BOT_OAUTH_ACCESS_TOKEN" required:"true"`
}

// New は設定されている環境変数をVariable構造体にセットして返します
func New() (*Variable, error) {
	var v Variable

	if err := envconfig.Process("", &v); err != nil {
		return nil, xerrors.Errorf("envconfig load error: %v", err)
	}

	return &v, nil
}
