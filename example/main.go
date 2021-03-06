package main

import (
	"log"
	"os"

	"github.com/0g3/slackbot"
)

func main() {
	whu := os.Getenv("SLACK_WEB_HOOK_URL")
	if whu == "" {
		log.Fatal("環境変数SLACK_WEB_HOOK_URLが設定されていません")
	}

	bot := slackbot.NewIncomingWebHookBot(whu)

	m := &slackbot.Message{
		Text: "送信フェイル！ フォールバック専用文字列と化した先輩!",
		Blocks: []slackbot.Block{
			slackbot.NewSectionBlock(&slackbot.SectionBlockDTO{
				Text: &slackbot.Text{
					Type: slackbot.TextTypeMkdwn,
					Text: "selectの文章です",
				},
				Accessory: slackbot.NewImageElement(&slackbot.ImageElementDTO{
					ImageURL: "http://trender-news.com/thumbnail/113635.jpg",
					AltText:  "可愛すぎて無理",
				}),
				Fields: []*slackbot.Text{
					{
						Type: slackbot.TextTypeMkdwn,
						Text: "フィールド1",
					},
					{
						Type: slackbot.TextTypeMkdwn,
						Text: "フィールド2",
					},
					{
						Type: slackbot.TextTypeMkdwn,
						Text: "フィールド3",
					},
					{
						Type: slackbot.TextTypeMkdwn,
						Text: "フィールド4",
					},
				},
			}),
			slackbot.NewDividerBlock(&slackbot.DividerBlockDTO{}),
			slackbot.NewImageBlock(&slackbot.ImageBlockDTO{
				ImageURL: "https://pbs.twimg.com/media/BtKE6XdCMAA0sYq.jpg",
				AltText:  "shinchoku",
				Title: &slackbot.Text{
					Type: slackbot.TextTypeMkdwn,
					Text: "image title!!!!!!!!!!!!",
				},
			}),
			slackbot.NewDividerBlock(&slackbot.DividerBlockDTO{}),
		},
	}
	if err := bot.Post(m); err != nil {
		log.Fatal(err)
	}

	if err := bot.PostTxt("これは *PostTxt関数* によってポストされた文章です。"); err != nil {
		log.Fatal(err)
	}
}
