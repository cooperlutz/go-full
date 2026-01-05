package entity

import (
	"github.com/cooperlutz/go-full/internal/pingpong/domain/constant"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/event"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/exception"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

// PingPongEntity represents the PingPong entity with its metadata and message.
// it represents the aggregate root for PingPong related operations.
type PingPongEntity struct {
	*baseentitee.EntityMetadata
	message string
	events  []interface{} // Domain events associated with the entity
}

// New creates a new PingPongEntity with the given message.
func New(msg string) (PingPongEntity, error) {
	ent := PingPongEntity{
		message:        msg,
		EntityMetadata: baseentitee.NewEntityMetadata(),
	}

	if err := ent.Validate(); err != nil {
		return PingPongEntity{}, err
	}

	ent.raiseDomainEvent(event.NewPingPongReceived(
		ent.GetIdString(),
		ent.GetMessage(),
	))

	return ent, nil
}

func (e PingPongEntity) GetDomainEvents() []interface{} {
	return e.events
}

func (e *PingPongEntity) raiseDomainEvent(event interface{}) {
	e.events = append(e.events, event)
}

// Validate checks if the PingPongEntity is valid.
func (e PingPongEntity) Validate() error {
	if e.message != constant.PingMessage && e.message != constant.PongMessage {
		return exception.ErrPingPongMsgValidation{}
	}

	return nil
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

// DetermineResponseMessage returns the appropriate response message baseentiteed on the current message.
func (e PingPongEntity) DetermineResponseMessage() string {
	if e.message == constant.PingMessage {
		return constant.PongFunMessage
	}

	if e.message == constant.PongMessage {
		return constant.PingFunMessage
	}

	return ""
}

// ListOfPingPongs represents a collection of PingPong entities.
type ListOfPingPongs struct {
	PingPongs []PingPongEntity
}

// MapToEntity accepts the raw values from the given parameters to construct a PingPongEntity.
// It should ONLY be used for reconstructing entities from stored data.
func MapToEntity(
	msg string,
	metadata *baseentitee.EntityMetadata,
) PingPongEntity {
	return PingPongEntity{
		message:        msg,
		EntityMetadata: metadata,
	}
}
