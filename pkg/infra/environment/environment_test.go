package environment

import (
	"os"
	"testing"

	"github.com/kichion/mischievous-slack-bot/pkg/infra/test"
)

func TestNew(t *testing.T) {
	os.Clearenv()
	os.Setenv("VERIFICATION_TOKEN", "test_VERIFICATION_TOKEN")

	v, err := New()
	if msg := test.Equal(err, nil, "New()"); msg != "" {
		t.Error(msg)
	}
	if msg := test.Equal(v.Slack.VerificationToken, "test_VERIFICATION_TOKEN", "VerificationToken"); msg != "" {
		t.Error(msg)
	}
}
