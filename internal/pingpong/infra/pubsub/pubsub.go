package pubsub

import (
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/repository"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Ensure EventProcessor implements the IPubSubEventProcessor interface.
var _ eeventdriven.IPubSubEventProcessor = (*PingPongPubSub)(nil)

type PingPongPubSub struct {
	*eeventdriven.BasePubSubProcessor
	repo repository.IPingPongRepository
}

func New(db *pgxpool.Pool, repo repository.IPingPongRepository) *PingPongPubSub {
	basePS := eeventdriven.New(db)
	ps := &PingPongPubSub{
		BasePubSubProcessor: basePS,
		repo:                repo,
	}

	return ps
}

func (pp *PingPongPubSub) RegisterHandlers() error {
	logger := watermill.NewStdLogger(false, false)

	subscriber, err := sql.NewSubscriber(
		sql.BeginnerFromPgx(pp.GetDB()),
		sql.SubscriberConfig{
			SchemaAdapter:    sql.DefaultPostgreSQLSchema{},
			OffsetsAdapter:   sql.DefaultPostgreSQLOffsetsAdapter{},
			InitializeSchema: true,
		},
		logger,
	)
	if err != nil {
		return err
	}

	publisher, err := sql.NewPublisher(
		sql.BeginnerFromPgx(pp.GetDB()),
		sql.PublisherConfig{
			SchemaAdapter: sql.DefaultPostgreSQLSchema{},
		},
		logger,
	)
	if err != nil {
		return err
	}

	pp.SetPublisher(publisher)

	router := pp.GetRouter()

	router.AddHandler(
		"_handler",
		"pingpong",
		subscriber,
		"pingpong",
		publisher,
		func(msg *message.Message) ([]*message.Message, error) {
			log.Printf("Processing message: %s, payload: %s", msg.UUID, string(msg.Payload))
			msg.Ack()
			return []*message.Message{msg}, nil
		},
	)

	return nil
}
