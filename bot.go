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
// ImageBlock
//==============================================================================
func NewImageBlock(i *ImageBlockDTO) Block {
	block := make(Block)
	block["type"] = BlockTypeImage

	block["image_url"] = i.ImageURL
	block["alt_text"] = i.AltText

	if i.Title != nil {
		i.Title.Type = TextTypePlainText // これしか許されない
		block["title"] = i.Title
	}
	if i.BlockID != "" {
		block["block_id"] = i.BlockID
	}

	return block
}

type ImageBlockDTO struct {
	ImageURL string
	AltText  string
	Title    *Text
	BlockID  string
}

//==============================================================================
// DividerBlock
//==============================================================================
func NewDividerBlock(d *DividerBlockDTO) Block {
	block := make(Block)
	block["type"] = BlockTypeDivider
	if d.BlockID != "" {
		block["block_id"] = d.BlockID
	}
	return block
}

type DividerBlockDTO struct {
	BlockID string
}

//==============================================================================
// SectionBlock
//==============================================================================
func NewSectionBlock(s *SectionBlockDTO) Block {
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

type SectionBlockDTO struct {
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

func NewImageElement(i *ImageElementDTO) BlockElement {
	be := make(BlockElement)
	be["type"] = BlockElementTypeImage
	be["image_url"] = i.ImageURL
	be["alt_text"] = i.AltText
	return be
}

type ImageElementDTO struct {
	ImageURL string
	AltText  string
}
