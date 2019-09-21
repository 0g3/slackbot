package slackbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

//==============================================================================
// Incomming web hook bot
//==============================================================================
type IncomingWebHookBot struct {
	webHookURL string
}

func NewIncomingWebHookBot(webHookURL string) *IncomingWebHookBot {
	return &IncomingWebHookBot{
		webHookURL: webHookURL,
	}
}

// Post はメッセージを送信します。
func (b *IncomingWebHookBot) Post(m *Message) error {
	byt, err := json.Marshal(m)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("[DEBUG] %s\n", string(byt))
	buf := bytes.NewBuffer(byt)
	req, err := http.NewRequest(
		"POST",
		b.webHookURL,
		buf,
	)
	if err != nil {
		return errors.New("failed to create request")
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

// PostTxt はテキストメッセージを送信します。
// mrkdwnはtrueです。
func (b *IncomingWebHookBot) PostTxt(txt string) error {
	return b.Post(&Message{
		Text:   txt,
		Mrkdwn: true,
	})
}

//==============================================================================
// Message
//==============================================================================
type Message struct {
	Text     string  `json:"text"`
	Blocks   []Block `json:"blocks"`
	ThreadTS string  `json:"thread_ts,omitempty"`
	Mrkdwn   bool    `json:"mrkdwn"`
}

type Block map[string]interface{}

type BlockType string

const (
	BlockTypeSection BlockType = "section"
	BlockTypeDivider BlockType = "divider"
	BlockTypeImage   BlockType = "image"
	BlockTypeActions BlockType = "actions"
	BlockTypeContext BlockType = "context"
	BlockTypeFile    BlockType = "file"
)

//==============================================================================
// Block: Divider
//==============================================================================
func NewDividerBlock(d *Divider) Block {
	block := make(Block)
	block["type"] = BlockTypeDivider
	if d.BlockID != "" {
		block["block_id"] = d.BlockID
	}
	return block
}

type Divider struct {
	BlockID string
}

//==============================================================================
// Block: Section
//==============================================================================
func NewSectionBlock(s *Section) Block {
	block := make(Block)
	block["type"] = BlockTypeSection

	if s.BlockID != "" {
		block["block_id"] = s.BlockID
	}

	block["text"] = s.Text

	if s.Accessory != nil {
		block["accessory"] = s.Accessory
	}

	if s.Fields != nil {
		block["fields"] = s.Fields
	}

	return block
}

type Section struct {
	BlockID   string
	Text      *Text
	Accessory BlockElement
	Fields    []*Text
}

type TextType string

const (
	TextTypeMkdwn     TextType = "mrkdwn"
	TextTypePlainText TextType = "plain_text"
)

type Text struct {
	Type     TextType `json:"type"`
	Text     string   `json:"text"`
	Emoji    *bool    `json:"emoji,omitempty"`
	Verbatim *bool    `json:"verbatim,omitempty"`
}

type BlockElementType string

// TODO: 実装する
const (
	BlockElementTypeImage               BlockElementType = "image"
	BlockElementTypeButton              BlockElementType = "button"
	BlockElementTypeStaticSelect        BlockElementType = "static_select"
	BlockElementTypeExternalSelect      BlockElementType = "external_select"
	BlockElementTypeUsersSelect         BlockElementType = "users_select"
	BlockElementTypeConversationsSelect BlockElementType = "conversations_select"
	BlockElementTypeChannelsSelect      BlockElementType = "channels_select"
	BlockElementTypeOverflow            BlockElementType = "overflow"
	BlockElementTypeDatepicker          BlockElementType = "datepicker"
)

type BlockElement map[string]interface{}

func NewBlockElementImage(i *ImageElement) BlockElement {
	be := make(BlockElement)
	be["type"] = BlockElementTypeImage
	be["image_url"] = i.ImageURL
	be["alt_text"] = i.AltText
	return be
}

type ImageElement struct {
	ImageURL string
	AltText  string
}
