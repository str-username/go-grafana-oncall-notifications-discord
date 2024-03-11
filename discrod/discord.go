package discrod

import (
	"bytes"
	"encoding/json"
)

type Message struct {
	Content string `json:"content"`
}

func (message *Message) Notify(text string) *bytes.Buffer {
	message.Content = text

	jsonData, err := json.Marshal(message)

	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(jsonData)
}
