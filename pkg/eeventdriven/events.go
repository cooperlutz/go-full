package eeventdriven

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

func EventPayloadToMessage(payload interface{}) (*message.Message, error) {
	marshaled, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	msg := message.NewMessage(watermill.NewUUID(), marshaled)
	return msg, nil
}
