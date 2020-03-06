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
	"golang.org/x/xerrors"
)

// CurryOrder はカレーのオーダーを受けるための振る舞いです
func CurryOrder(e *slackevents.AppMentionEvent, v *environment.Variable) (events.APIGatewayProxyResponse, error) {
	headerText := slack.NewTextBlockObject("mrkdwn", "*カレーはいかがっすか?*", false, false)
	divider := slack.NewDividerBlock()

	api := slack.New(v.Slack.OAuthAccessToken)

	if _, _, err := api.PostMessage(
		e.Channel,
		slack.MsgOptionBlocks(
			slack.NewSectionBlock(headerText, nil, nil),
			divider,
			createCurrySection("ビーフカレー", 400, "シンプルなカレーライス", "vote_beef"),
			createCurrySection("ビーフカレー大盛り", 500, "大食らいをも満たすシンプルなカレーライス", "vote_big_beef"),
			createCurrySection("カツカレー", 500, "物足りなさを感じさせないわがままなカレーライス", "vote_cutlet"),
			createCurrySection("カツカレー大盛り", 600, "コレでだめなら自分で作れなカレーライス", "vote_big_cutlet"),
			divider,
		),
	); err != nil {
		printErr := xerrors.Errorf("Talk error: %v", err)
		log.Print(printErr)
		return responce.NewGateway(http.StatusBadRequest), printErr
	}

	return responce.NewGateway(http.StatusOK), nil
}

func createCurrySection(name string, amount int, desc string, btnVal string) *slack.SectionBlock {
	txt := slack.NewTextBlockObject(
		"mrkdwn",
		fmt.Sprintf(":dollar: ¥%d【:curry: *%s*】\n%s", name, amount, desc),
		false,
		false,
	)
	btnTxt := slack.NewTextBlockObject("plain_text", "追加", true, false)
	btn := slack.NewButtonBlockElement("", btnVal, btnTxt)
	return slack.NewSectionBlock(txt, nil, slack.NewAccessory(btn))
}
