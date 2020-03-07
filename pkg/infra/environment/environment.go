package environment

import (
	"golang.org/x/xerrors"

	"github.com/kelseyhightower/envconfig"
)

// Variable は環境変数の情報を保持する構造体です
type Variable struct {
	Slack      Slack
	A3RT       A3RT
	TalkMaster TalkMaster
}

// Slack はSlackにおける情報を保持する構造体です
type Slack struct {
	VerificationToken string `envconfig:"VERIFICATION_TOKEN" required:"true"`
	OAuthAccessToken  string `envconfig:"BOT_OAUTH_ACCESS_TOKEN" required:"true"`
	SigningSecret     string `envconfig:"SIGNING_SECRET" required:"true"`
	BotMention        string `envconfig:"BOT_MENTION" required:"true"`
}

// A3RT はA3RTのAPIにおける情報を保持する構造体です
type A3RT struct {
	TalkAPIKey         string `envconfig:"TALK_API_KEY" required:"true"`
	ProofreadingAPIKey string `envconfig:"PROOFREADING_API_KEY" required:"true"`
}

// TalkMaster は会話用の言語マスタにおける情報を保持する構造体です
type TalkMaster struct {
	SpreadsheetID string `envconfig:"TALK_MASTER_SPREADSHEET_ID" required:"true"`
	S3Storage     string `envconfig:"TALK_MASTER_SERCRET_S3" required:"true"`
	AWSRegion     string `envconfig:"TALK_MASTER_SERCRET_REGION" required:"true"`
}

// New は設定されている環境変数をVariable構造体にセットして返します
func New() (*Variable, error) {
	var v Variable

	if err := envconfig.Process("", &v); err != nil {
		return nil, xerrors.Errorf("envconfig load error: %v", err)
	}

	return &v, nil
}
