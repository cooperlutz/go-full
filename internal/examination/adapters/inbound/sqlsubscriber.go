package inbound

import (
	"context"

	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type SqlSubscriberAdapter struct {
	subscriber *sql.Subscriber
	router     *message.Router
}

func NewSqlSubscriberAdapter(db deebee.IDatabase) (SqlSubscriberAdapter, error) {
	router, err := eeventdriven.InitTracedRouter()
	if err != nil {
		return SqlSubscriberAdapter{}, err
	}

	subscriber, err := sql.NewSubscriber(
		sql.BeginnerFromPgx(db),
		sql.SubscriberConfig{
			SchemaAdapter:    sql.DefaultPostgreSQLSchema{},
			OffsetsAdapter:   sql.DefaultPostgreSQLOffsetsAdapter{},
			InitializeSchema: true,
		},
		nil,
	)
	if err != nil {
		return SqlSubscriberAdapter{}, err
	}

	return SqlSubscriberAdapter{
		subscriber: subscriber,
		router:     router,
	}, nil
}

func (s SqlSubscriberAdapter) GetRouter() *message.Router {
	return s.router
}

func (s SqlSubscriberAdapter) Start() {
	ctx := context.Background()

	err := s.router.Run(ctx)
	if err != nil {
		panic(err)
	}
}

func (s SqlSubscriberAdapter) RegisterEventHandler(handler message.NoPublishHandlerFunc) {
	s.router.AddConsumerHandler(
		"examination_examsubmitted_handler",
		"examination.examsubmitted",
		s.subscriber,
		handler,
	)
	s.router.AddConsumerHandler(
		"examination_examstarted_handler",
		"examination.examstarted",
		s.subscriber,
		handler,
	)
}
