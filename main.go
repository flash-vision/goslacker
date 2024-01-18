package goslacker

import (
    "fmt"
    
)

// SlackMessage struct with methods for chaining
type SlackMessage struct {
    Blocks []Block `json:"blocks"`
}

func NewSlackMessage() *SlackMessage {
    return &SlackMessage{}
}

func (m *SlackMessage) AddBlock(block *Block) *SlackMessage {
    m.Blocks = append(m.Blocks, *block)
    return m
}

func (m *SlackMessage) Validate() error {
    // Add validation logic here
    return nil
}

func (m *SlackMessage) Print() {
    fmt.Println(m)
}

// Block struct with methods and builder pattern
type Block struct {
    Type      string       `json:"type"`
    Text      *TextObject  `json:"text,omitempty"`
    Elements  []TextObject `json:"elements,omitempty"`
    Title     *TextObject  `json:"title,omitempty"`
    ImageURL  string       `json:"image_url,omitempty"`
    AltText   string       `json:"alt_text,omitempty"`
    Accessory *Accessory   `json:"accessory,omitempty"`
}

func NewBlock() *Block {
    return &Block{}
}

func (b *Block) AddTextObject(textObject *TextObject) *Block {
    b.Elements = append(b.Elements, *textObject)
    return b
}

func (b *Block) AddAccessory(accessory *Accessory) *Block {
    b.Accessory = accessory
    return b
}

func (b *Block) SetTitle(title *TextObject) *Block {
    b.Title = title
    return b
}

func (b *Block) SetImageURL(imageURL string) *Block {
    b.ImageURL = imageURL
    return b
}

func (b *Block) SetAltText(altText string) *Block {
    b.AltText = altText
    return b
}

// TextObject struct
type TextObject struct {
    Type  string `json:"type"`
    Text  string `json:"text"`
    Emoji bool   `json:"emoji,omitempty"`
}

func NewTextObject(textType, text string, emoji bool) *TextObject {
    return &TextObject{
        Type:  textType,
        Text:  text,
        Emoji: emoji,
    }
}

// Accessory struct
type Accessory struct {
    Type     string      `json:"type"`
    ImageURL string      `json:"image_url,omitempty"`
    AltText  string      `json:"alt_text,omitempty"`
    Text     *TextObject `json:"text,omitempty"`
    Value    string      `json:"value,omitempty"`
    URL      string      `json:"url,omitempty"`
    ActionID string      `json:"action_id,omitempty"`
}

func NewAccessory(accessoryType string, text *TextObject) *Accessory {
    return &Accessory{
        Type: accessoryType,
        Text: text,
    }
}
