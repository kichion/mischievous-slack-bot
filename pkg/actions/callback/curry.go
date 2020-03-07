package callback

import (
	"strconv"
	"strings"

	"github.com/slack-go/slack"
)

const (
	noVote     = "注文なし"
	voteSuffix = " 件"
)

// CurrySelect はカレー選択のコールバック結果を含めたBlockを返します
func CurrySelect(payload *slack.InteractionCallback) []slack.Block {
	action := payload.ActionCallback.BlockActions[0]
	blocks := payload.Message.Blocks.BlockSet
	userName := payload.User.Name

	for i, b := range blocks {
		switch block := b.(type) {
		case *slack.ContextBlock:
			if block.BlockID != action.ActionID {
				continue
			}

			blocks[i] = build(block, userName)
		}
	}
	return blocks
}

func build(block *slack.ContextBlock, userName string) *slack.ContextBlock {
	base, deleted := createBase(block, userName)
	if !deleted {
		base = append(base, slack.TextBlockObject{
			Type:     "plain_text",
			Text:     userName,
			Emoji:    false,
			Verbatim: false,
		})
	}
	if len(base) == 0 {
		base = append(base, slack.TextBlockObject{
			Type:     "plain_text",
			Text:     noVote,
			Emoji:    false,
			Verbatim: false,
		})
	} else {
		base = append(base, slack.TextBlockObject{
			Type:     "plain_text",
			Text:     strconv.Itoa(len(base)) + voteSuffix,
			Emoji:    false,
			Verbatim: false,
		})
	}
	return slack.NewContextBlock(block.BlockID, base...)
}

func createBase(block *slack.ContextBlock, userName string) ([]slack.MixedElement, bool) {
	userDeleted := false
	result := []slack.MixedElement{}
	for _, e := range block.ContextElements.Elements {
		switch elem := e.(type) {
		case *slack.TextBlockObject:
			if elem.Text == noVote || strings.Contains(elem.Text, voteSuffix) {
				continue
			}
			if elem.Text == userName {
				userDeleted = true
				continue
			}
			result = append(result, elem)
		}
	}
	return result, userDeleted
}
