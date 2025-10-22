package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/constant"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/exception"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

type PingPongMetadata struct {
	PingPongID uuid.UUID  `json:"pingpong_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Deleted    bool       `json:"deleted"`
}

type PingPongEntity struct {
	Message string `json:"message"`
	*PingPongMetadata
}

func New(msg string) (*PingPongEntity, error) {
	ent := &PingPongEntity{
		Message: msg,
		PingPongMetadata: &PingPongMetadata{
			PingPongID: uuid.New(),
			CreatedAt:  utilitee.RightNow(),
			UpdatedAt:  utilitee.RightNow(),
			DeletedAt:  nil,
			Deleted:    false,
		},
	}

	if err := ent.Validate(); err != nil {
		return nil, err
	}

	return ent, nil
}

type ListOfPingPongs struct {
	PingPongs []PingPongEntity
}

func (e *PingPongEntity) Validate() error {
	if e.Message != constant.PingMessage && e.Message != constant.PongMessage {
		return exception.ErrPingPongMsgValidation{}
	}

	return nil
}

func (e *PingPongEntity) Valid() bool {
	return e.Validate() == nil
}

func (e *PingPongEntity) GetMessage() string {
	return e.Message
}

func (e *PingPongEntity) DetermineResponseMessage() string {
	if e.Message == constant.PingMessage {
		return constant.PongFunMessage
	}

	if e.Message == constant.PongMessage {
		return constant.PingFunMessage
	}

	return ""
}
