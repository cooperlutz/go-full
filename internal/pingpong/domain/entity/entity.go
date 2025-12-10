package entity

import (
	"github.com/cooperlutz/go-full/internal/pingpong/domain/constant"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/exception"
	"github.com/cooperlutz/go-full/pkg/base"
)

type ListOfPingPongs struct {
	PingPongs []PingPongEntity
}

type PingPongEntity struct {
	base.EntityMetadata
	message string
}

func New(msg string) (PingPongEntity, error) {
	ent := PingPongEntity{
		message:        msg,
		EntityMetadata: base.NewEntityMetadata(),
	}

	if err := ent.Validate(); err != nil {
		return PingPongEntity{}, err
	}

	return ent, nil
}

// Validate checks if the PingPongEntity is valid.
func (e PingPongEntity) Validate() error {
	if e.message != constant.PingMessage && e.message != constant.PongMessage {
		return exception.ErrPingPongMsgValidation{}
	}

	return nil
}

// Valid returns true if the PingPongEntity is valid.
func (e PingPongEntity) Valid() bool {
	return e.Validate() == nil
}

// GetMessage returns the message of the PingPongEntity.
func (e PingPongEntity) GetMessage() string {
	return e.message
}

// SetMessage sets the message of the PingPongEntity and marks it as updated.
func (e *PingPongEntity) SetMessage(msg string) {
	e.message = msg
	e.MarkUpdated()
}

// DetermineResponseMessage returns the appropriate response message based on the current message.
func (e PingPongEntity) DetermineResponseMessage() string {
	if e.message == constant.PingMessage {
		return constant.PongFunMessage
	}

	if e.message == constant.PongMessage {
		return constant.PingFunMessage
	}

	return ""
}

func MapToEntity(
	msg string,
	metadata base.EntityMetadata,
) PingPongEntity {
	return PingPongEntity{
		message:        msg,
		EntityMetadata: metadata,
	}
}
