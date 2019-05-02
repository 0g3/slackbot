package slackbot

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type Field struct {
	// valueの上に太字で見出し
	Title string `json:"title"`
	Value string `json:"value"`
	// true: 横に並べられるなら横に並べる
	Short bool `json:"short"`
}

type Attachment struct {
	// 要約メッセージ。通知やモバイル端末での表示に使われる。
	Fallback string `json:"fallback"`

	// アタッチメントブロック(左に線が引いてある部分)の上に表示する文字。
	Pretext string `json:"pretext"`

	AuthorName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	AuthorIcon string `json:"author_icon"`

	// アタッチメントボックスの先頭に太字で表示される文字
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`

	// attachmentの本文
	Text string `json:"text"`

	Fields []*Field `json:"fields"`

	// ラインの色。Slackで用意されているgood、warning、dangerの３つを指定するか、カラーコードを指定する。
	Color string `json:"color"`

	// attachmentに埋め込む画像URL
	ImageURL string `json:"image_url"`
	// attachmentに埋め込むサムネイル画像URL。image_urlとは画像の表示のされ方が違う。
	// ImageURLが設定されていないときのみ有効になることに注意！
	ThumbURL string `json:"thumb_url"`

	Footer     string `json:"footer"`
	FooterIcon string `json:"footer_icon"`

	// attachmentに付加する情報に対するタイムスタンプ(UNIXタイムスタンプ)
	TS string `json:"ts"`
}

type PostMessage struct {
	Text        string        `json:"text"`
	Attachments []*Attachment `json:"attachments"`
}

type IncomingWebHookBot struct {
	WebHookURL string
}

func (b *IncomingWebHookBot) Post(p *PostMessage) error {
	byt, err := json.Marshal(p)
	if err != nil {
		return errors.WithStack(err)
	}
	buf := bytes.NewBuffer(byt)
	req, err := http.NewRequest(
		"POST",
		b.WebHookURL,
		buf,
	)
	if err != nil {
		return errors.New("リクエストの生成に失敗しました")
	}

	req.Header.Set("Content-Type", "application/json")

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	return nil
}

func (b *IncomingWebHookBot) PostTxt(txt string) error {
	return b.Post(&PostMessage{Text: txt})
}
