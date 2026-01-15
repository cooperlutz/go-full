package adapters

import (
	"github.com/cooperlutz/go-full/internal/pingpong/domain/repository"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

// Ensure EventProcessor implements the IPubSubEventProcessor interface.
var _ eeventdriven.IPubSubEventProcessor = (*hourPubSub)(nil)

// hourPubSub handles Pub/Sub events specific to the hour module.
type hourPubSub struct {
	*eeventdriven.BasePgsqlPubSubProcessor
	repo repository.IPingPongRepository
}

// New - Creates a new instance of hourPubSub with the provided database connection and repository.
func NewHourPubSub(db deebee.IDatabase, repo repository.IPingPongRepository) (*hourPubSub, error) {
	basePS, err := eeventdriven.NewPubSub(db)
	if err != nil {
		return nil, err
	}

	ps := &hourPubSub{
		BasePgsqlPubSubProcessor: basePS,
		repo:                     repo,
	}

	return ps, nil
}

// // RegisterHandlers overrides the base method to register hour-specific event handlers.
// func (pp *hourPubSub) RegisterSubscriberHandler(handler *message.Handler) error {
// 	router := pp.GetRouter()

// 	return nil
// }
