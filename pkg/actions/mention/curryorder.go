package mention

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kichion/mischievous-slack-bot/pkg/domain/valueobject/responce"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/environment"
	"github.com/kichion/mischievous-slack-bot/pkg/infra/order"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"golang.org/x/xerrors"
)

// CurryOrder はカレーのオーダーを受けるための振る舞いです
func CurryOrder(e *slackevents.AppMentionEvent, v *environment.Variable) (events.APIGatewayProxyResponse, error) {
	headerText := slack.NewTextBlockObject("mrkdwn", "*カレーはいかがっすか?*", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	divider := slack.NewDividerBlock()

	client, err := order.NewClient(&v.OrderMaster)
	if err != nil {
		printErr := xerrors.Errorf("Order error: %v", err)
		log.Print(printErr)
		return responce.NewGateway(http.StatusBadRequest), printErr
	}
	items, err := client.OrderList(context.Background())
	if err != nil {
		printErr := xerrors.Errorf("Order error: %v", err)
		log.Print(printErr)
		return responce.NewGateway(http.StatusBadRequest), printErr
	}

	blocks := []slack.Block{
		headerSection,
		divider,
	}
	for _, item := range items {
		blocks = append(
			blocks,
			createSection(item.Name, item.Price, item.Icon, item.Desc, item.UUID),
			createVoteContext(item.UUID),
		)
	}
	blocks = append(blocks, divider)

	api := slack.New(v.Slack.OAuthAccessToken)

	if _, _, err = api.PostMessage(
		e.Channel,
		slack.MsgOptionBlocks(blocks...),
	); err != nil {
		printErr := xerrors.Errorf("Order error: %v", err)
		log.Print(printErr)
		return responce.NewGateway(http.StatusBadRequest), printErr
	}

	return responce.NewGateway(http.StatusOK), nil
}

func createSection(name string, price int, icon string, desc string, uuid string) *slack.SectionBlock {
	txt := slack.NewTextBlockObject(
		"mrkdwn",
		fmt.Sprintf(":dollar: ¥%d【:%s: *%s* 】\n%s", price, icon, name, desc),
		false,
		false,
	)
	btnTxt := slack.NewTextBlockObject("plain_text", "追加", true, false)
	btn := slack.NewButtonBlockElement(uuid, uuid, btnTxt)

	return slack.NewSectionBlock(txt, nil, slack.NewAccessory(btn))
}

func createVoteContext(uuid string) *slack.ContextBlock {
	contextTxt := slack.NewTextBlockObject("plain_text", "注文なし", true, false)
	context := slack.NewContextBlock(uuid, []slack.MixedElement{contextTxt}...)

	return context
}
