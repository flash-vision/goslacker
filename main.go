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

func AddBlock(message SlackMessage, block Block) SlackMessage {
	message.Blocks = append(message.Blocks, block)
	return message
}

func AddTextObject(block Block, textObject TextObject) Block {
	block.Elements = append(block.Elements, textObject)
	return block
}

func AddAccessory(block Block, accessory Accessory) Block {
	block.Accessory = &accessory
	return block
}

func AddElement(block Block, element TextObject) Block {
	block.Elements = append(block.Elements, element)
	return block
}

func SetTitle(block Block, title TextObject) Block {
	block.Title = &title
	return block
}

func SetImageURL(block Block, imageURL string) Block {
	block.ImageURL = imageURL
	return block
}

func SetAltText(block Block, altText string) Block {
	block.AltText = altText
	return block
}

func SetText(textObject TextObject, text string) TextObject {
	textObject.Text = text
	return textObject
}

func PrintMessage(message SlackMessage) {
	fmt.Println(message)
}

