package actions

import "github.com/slack-go/slack/slackevents"

// IsURLVerification はEventがURL検証のものかを検査します
func IsURLVerification(event slackevents.EventsAPIEvent) bool {
	return event.Type == slackevents.URLVerification
}

// IsUndefined はEventがアプリケーション上未定義のものかを検査します
func IsUndefined(event slackevents.EventsAPIEvent) bool {
	return event.Type != slackevents.CallbackEvent
}
