package goslacker
import (
	"fmt"
)

type SlackMessage struct {
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type      string       `json:"type"`
	Text      *TextObject  `json:"text,omitempty"`
	Elements  []TextObject `json:"elements,omitempty"`
	Title     *TextObject  `json:"title,omitempty"`
	ImageURL  string       `json:"image_url,omitempty"`
	AltText   string       `json:"alt_text,omitempty"`
	Accessory *Accessory   `json:"accessory,omitempty"`
}

type TextObject struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"`
}

type Accessory struct {
	Type     string      `json:"type"`
	ImageURL string      `json:"image_url,omitempty"`
	AltText  string      `json:"alt_text,omitempty"`
	Text     *TextObject `json:"text,omitempty"`
	Value    string      `json:"value,omitempty"`
	URL      string      `json:"url,omitempty"`
	ActionID string      `json:"action_id,omitempty"`
}

/*
  This contains an interface and function for creating a
  slack message construct.  The interface is used to
  create a message that can be sent to slack.  
  The NewSlackMessage function is used to create a new message.
  The AddBlock function is used to add a block to the message, and
  takes a a SlackMessage and a Block as arguments.
  The AddTextObject function is used to add a text object to a block,
  and takes a Block and a TextObject as arguments.
  The AddAccessory function is used to add an accessory to a block,
  and takes a Block and an Accessory as arguments.
  The AddElement function is used to add an element to a block,
  and takes a Block and a TextObject as arguments.
  The SetTitle function is used to set the title of a block,
  and takes a Block and a TextObject as arguments.
  The SetImageURL function is used to set the image URL of a block,
  and takes a Block and a string as arguments.
  The SetAltText function is used to set the alt text of a block,
  and takes a Block and a string as arguments.
  The SetText function is used to set the text of a text object,
  and takes a TextObject and a string as arguments.
*/

type SlackMessageInterface interface {
	NewSlackMessage() SlackMessage
	AddBlock(SlackMessage, Block) SlackMessage
	AddTextObject(Block, TextObject) Block
	AddAccessory(Block, Accessory) Block
	AddElement(Block, TextObject) Block
	SetTitle(Block, TextObject) Block
	SetImageURL(Block, string) Block
	SetAltText(Block, string) Block
	SetText(TextObject, string) TextObject
	PrintMessage(SlackMessage)
}

func NewSlackMessage() SlackMessage {
	return SlackMessage{}
}
func (m *SlackMessage) AddBlock(block Block) *SlackMessage {
	m.Blocks = append(m.Blocks, block)
	return m
}

func (m *SlackMessage) AddTextObject(block Block, textObject TextObject) *SlackMessage {
	block.Elements = append(block.Elements, textObject)
	return m
}

func (m *SlackMessage) AddAccessory(block Block, accessory Accessory) *SlackMessage {
	block.Accessory = &accessory
	return m
}

func (m *SlackMessage) AddElement(block Block, element TextObject) *SlackMessage {
	block.Elements = append(block.Elements, element)
	return m
}

func (m *SlackMessage) SetTitle(block Block, title TextObject) *SlackMessage {
	block.Title = &title
	return m
}

func (m *SlackMessage) SetImageURL(block Block, imageURL string) *SlackMessage {
	block.ImageURL = imageURL
	return m
}


func (m *SlackMessage) SetAltText(block Block, altText string) *SlackMessage {
	block.AltText = altText
	return m
}

func (m *SlackMessage) SetText(textObject TextObject, text string) *SlackMessage {
	textObject.Text = text
	return m
}

func (m *SlackMessage) PrintMessage() {
	fmt.Println(m)
}
