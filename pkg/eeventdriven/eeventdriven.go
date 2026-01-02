package eeventdriven

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IPubSubEventProcessor interface {
	EmitEvent(topic string, payload interface{}) error
	Run(wg *sync.WaitGroup)
	RegisterHandlers() error
}

type BasePubSubProcessor struct {
	db        *pgxpool.Pool
	router    *message.Router
	publisher *sql.Publisher
}

func New(db *pgxpool.Pool) *BasePubSubProcessor {
	logger := watermill.NewStdLogger(false, false)

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	ps := &BasePubSubProcessor{
		db:     db,
		router: router,
	}

	return ps
}

func (bps *BasePubSubProcessor) GetDB() *pgxpool.Pool {
	return bps.db
}

func (bps *BasePubSubProcessor) SetPublisher(publisher *sql.Publisher) {
	bps.publisher = publisher
}

func (bps *BasePubSubProcessor) EmitEvent(topic string, structPayload interface{}) error {
	marshalled, err := json.Marshal(structPayload)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), marshalled)
	if err := bps.publisher.Publish(topic, msg); err != nil {
		return err
	}

	return nil
}

func (bps *BasePubSubProcessor) GetRouter() *message.Router {
	return bps.router
}

func (bps *BasePubSubProcessor) RegisterHandlers() error {
	return errors.New("not implemented")
}

func (bps *BasePubSubProcessor) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	go func() {
		if err := bps.router.Run(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
}
