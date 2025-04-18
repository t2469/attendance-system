package services

import (
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/t2469/attendance-system.git/config"
	"log"
)

var LineBot *messaging_api.MessagingApiAPI

func InitLineBot(cfg *config.Config) {
	bot, err := messaging_api.NewMessagingApiAPI(cfg.LineChannelToken)
	if err != nil {
		log.Fatal(err)
	}
	LineBot = bot
}

func GetLineUserId(source webhook.SourceInterface) (string, bool) {
	switch s := source.(type) {
	case webhook.UserSource:
		return s.UserId, true
	case *webhook.UserSource:
		return s.UserId, true
	}
	return "", false
}

// Reply replyTokenを用いて、ユーザーへメッセージを送信 (LINE連携時などに使用)
func Reply(replyToken string, msg string) {
	_, err := LineBot.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: msg,
			},
		},
	})
	if err != nil {
		log.Printf(err.Error())
	}
}

// SendMessage UserIdが設定されている前提 (LINE連携済みユーザーへメッセージを送信する際に使用)
func SendMessage(UserId string, msg string) error {
	_, err := LineBot.PushMessage(
		&messaging_api.PushMessageRequest{
			To: UserId,
			Messages: []messaging_api.MessageInterface{
				&messaging_api.TextMessage{Text: msg},
			},
		}, "",
	)
	if err != nil {
		log.Printf(err.Error())
	}
	return err
}
