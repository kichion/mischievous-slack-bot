package environment

import (
	"golang.org/x/xerrors"

	"github.com/kelseyhightower/envconfig"
)

// Variable は環境変数の情報を保持する構造体です
type Variable struct {
	Slack Slack
	A3RT  A3RT
}

// Slack はSlackにおける情報を保持する構造体です
type Slack struct {
	VerificationToken string `envconfig:"VERIFICATION_TOKEN" required:"true"`
	OAuthAccessToken  string `envconfig:"BOT_OAUTH_ACCESS_TOKEN" required:"true"`
	SigningSecret     string `envconfig:"SIGNING_SECRET" required:"true"`
}

// A3RT はA3RTのAPIにおける情報を保持する構造体です
type A3RT struct {
	TalkAPIKey string `envconfig:"TALK_API_KEY" required:"true"`
	BaseURL    string `envconfig:"A3RT_BASE_URL" required:"true"`
}

// New は設定されている環境変数をVariable構造体にセットして返します
func New() (*Variable, error) {
	var v Variable

	if err := envconfig.Process("", &v); err != nil {
		return nil, xerrors.Errorf("envconfig load error: %v", err)
	}

	return &v, nil
}
