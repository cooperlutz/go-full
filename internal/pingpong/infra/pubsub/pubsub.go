package pubsub

import (
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/repository"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

// Ensure EventProcessor implements the IPubSubEventProcessor interface.
var _ eeventdriven.IPubSubEventProcessor = (*PingPongPubSub)(nil)

// PingPongPubSub handles Pub/Sub events specific to the PingPong module.
type PingPongPubSub struct {
	*eeventdriven.BasePgsqlPubSubProcessor
	repo repository.IPingPongRepository
}

// New - Creates a new instance of PingPongPubSub with the provided database connection and repository.
func New(db deebee.IDatabase, repo repository.IPingPongRepository) (*PingPongPubSub, error) {
	basePS, err := eeventdriven.New(db)
	if err != nil {
		return nil, err
	}

	ps := &PingPongPubSub{
		BasePgsqlPubSubProcessor: basePS,
		repo:                     repo,
	}

	return ps, nil
}

// RegisterHandlers overrides the base method to register PingPong-specific event handlers.
func (pp *PingPongPubSub) RegisterSubscriberHandlers() error {
	router := pp.GetRouter()

	router.AddHandler(
		"handler",
		"pingpong",
		pp.GetSubscriber(),
		"pingpong",
		pp.GetPublisher(),
		func(msg *message.Message) ([]*message.Message, error) {
			msg.Ack()

			return []*message.Message{msg}, nil
		},
	)

	return nil
}
