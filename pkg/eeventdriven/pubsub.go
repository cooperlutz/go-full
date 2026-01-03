package eeventdriven

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"

	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/pkg/deebee"
)

type IPubSubEventProcessor interface {
	EmitEvent(topic string, payload interface{}) error
	Run(wg *sync.WaitGroup)
	RegisterSubscriberHandlers() error
}

// BasePgsqlPubSubProcessor provides foundational Pub/Sub capabilities using Watermill and PostgreSQL.
type BasePgsqlPubSubProcessor struct {
	Config     config.Config
	db         deebee.IDatabase
	router     *message.Router
	publisher  *sql.Publisher
	subscriber *sql.Subscriber
}

// New initializes a new BasePgsqlPubSubProcessor with the given PostgreSQL connection pool.
func New(db deebee.IDatabase) (*BasePgsqlPubSubProcessor, error) {
	logger := watermill.NewStdLogger(false, false)

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddMiddleware(middleware.Recoverer)

	ps := &BasePgsqlPubSubProcessor{
		db:     db,
		router: router,
	}

	subscriber, err := sql.NewSubscriber(
		sql.BeginnerFromPgx(ps.GetDB()),
		sql.SubscriberConfig{
			SchemaAdapter:    sql.DefaultPostgreSQLSchema{},
			OffsetsAdapter:   sql.DefaultPostgreSQLOffsetsAdapter{},
			InitializeSchema: true,
		},
		logger,
	)
	if err != nil {
		return nil, err
	}

	ps.SetSubscriber(subscriber)

	publisher, err := sql.NewPublisher(
		sql.BeginnerFromPgx(ps.GetDB()),
		sql.PublisherConfig{
			SchemaAdapter: sql.DefaultPostgreSQLSchema{},
		},
		logger,
	)
	if err != nil {
		return nil, err
	}

	ps.SetPublisher(publisher)

	return ps, nil
}

// GetDB returns the PostgreSQL connection pool.
func (bps BasePgsqlPubSubProcessor) GetDB() deebee.IDatabase {
	return bps.db
}

// SetPublisher sets the SQL publisher for the Pub/Sub processor.
func (bps *BasePgsqlPubSubProcessor) SetPublisher(publisher *sql.Publisher) {
	bps.publisher = publisher
}

// GetPublisher returns the SQL publisher for the Pub/Sub processor.
func (bps BasePgsqlPubSubProcessor) GetPublisher() *sql.Publisher {
	return bps.publisher
}

// GetSubscriber returns the SQL subscriber for the Pub/Sub processor.
func (bps BasePgsqlPubSubProcessor) GetSubscriber() *sql.Subscriber {
	return bps.subscriber
}

// SetSubscriber sets the SQL subscriber for the Pub/Sub processor.
func (bps *BasePgsqlPubSubProcessor) SetSubscriber(subscriber *sql.Subscriber) {
	bps.subscriber = subscriber
}

// EmitEvent publishes an event with the given topic and payload.
func (bps *BasePgsqlPubSubProcessor) EmitEvent(topic string, payload interface{}) error {
	marshaled, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), marshaled)
	if err := bps.publisher.Publish(topic, msg); err != nil {
		return err
	}

	return nil
}

// GetRouter returns the message router.
func (bps BasePgsqlPubSubProcessor) GetRouter() *message.Router {
	return bps.router
}

// RegisterSubscriberHandlers registers the event handlers.
// By default, it returns an error indicating that handlers are not implemented.
func (bps *BasePgsqlPubSubProcessor) RegisterSubscriberHandlers() error {
	return &ErrPubSubHandlersNotImplemented{}
}

// Run starts the Pub/Sub processor's router.
func (bps *BasePgsqlPubSubProcessor) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		if err := bps.router.Run(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
}
