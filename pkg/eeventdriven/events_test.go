package eeventdriven

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventPayloadToMessage(t *testing.T) {
	t.Parallel()

	type SampleStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	payload := SampleStruct{
		Field1: "test",
		Field2: 100,
	}

	msg, err := EventPayloadToMessage(payload)
	assert.NoError(t, err)
	assert.NotNil(t, msg)

	var decodedPayload SampleStruct
	err = json.Unmarshal(msg.Payload, &decodedPayload)
	assert.NoError(t, err)
	assert.Equal(t, payload, decodedPayload)
}

func TestEventPayloadToMessage_JsonDecodeError(t *testing.T) {
	t.Parallel()

	type SampleStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}
	payload := map[string]interface{}{
		"foo": make(chan int), // Json cannot marshal channels
	}

	got, err := EventPayloadToMessage(payload)
	assert.Error(t, err)
	assert.Nil(t, got)

}
