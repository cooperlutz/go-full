package outbound

import (
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

// SqlPublisherAdapter provides an adapter to publish messages to a SQL database.
type SqlPublisherAdapter struct {
	Publisher eeventdriven.IPubSubEventProcessor
}

// NewSqlPublisherAdapter creates a new SqlPublisherAdapter using the provided database connection.
func NewSqlPublisherAdapter(pubSub eeventdriven.IPubSubEventProcessor) SqlPublisherAdapter {
	return SqlPublisherAdapter{
		Publisher: pubSub,
	}
}
