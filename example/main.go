package main

import (
	"github.com/0g3/slackbot"
	"log"
	"os"
)

func main() {
	whu := os.Getenv("SLACK_WEB_HOOK_URL")
	if whu == "" {
		log.Fatal("環境変数SLACK_WEB_HOOK_URLが設定されていません")
	}
	bot := slackbot.IncomingWebHookBot{WebHookURL: whu}
	p := &slackbot.PostMessage{
		Text: "これはPostMessageのTextです。",
		Attachments: []*slackbot.Attachment{
			{
				Pretext:    "これはプレテキストです。",
				AuthorName: "著者名: BanG Dream!",
				AuthorLink: "https://bang-dream.com/",
				AuthorIcon: "https://pbs.twimg.com/profile_images/1106571190863196160/iKEpq-qE_400x400.png",
				Title:      "タイトル: ガールズバンドパーティ！",
				TitleLink:  "https://bang-dream.bushimo.jp/",
				Text:       "キラキラでドキドキなテキストです。",
				Fields: []*slackbot.Field{
					{
						Title: "戸山香澄",
						Value: "愛美",
						Short: true,
					},
					{
						Title: "花園たえ",
						Value: "大塚紗英",
						Short: true,
					},
					{
						Title: "牛込りみ",
						Value: "西本りみ",
						Short: true,
					},
					{
						Title: "山吹沙綾",
						Value: "大橋彩香",
						Short: true,
					},
					{
						Title: "市ヶ谷有咲",
						Value: "伊藤彩沙",
						Short: true,
					},
				},
				Color: "#f44336",
				//ImageURL: "https://s3-ap-northeast-1.amazonaws.com/bang-dream.bushimo.jp/wordpress/wp-content/uploads/2019/03/190316_press_2ndAnniversary.jpg",
				ThumbURL:   "https://www.c-labo.jp/wordpress/wp-content/uploads/2018/07/40103bd2dd1917141e8054f34db1e618.png",
				Footer:     "footer: Poppin'Party",
				FooterIcon: "http://www.neowing.co.jp/pictures/l/13/07/BRMM-10171.jpg",
				TS:         "1556636400",
			},
		},
	}

	if err := bot.Post(p); err != nil {
		log.Fatal(err)
	}

	if err := bot.PostTxt("これはPostTxtによってポストされた文章です。"); err != nil {
		log.Fatal(err)
	}
}
