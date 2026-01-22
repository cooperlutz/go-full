package outbound

import (
	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"

	"github.com/cooperlutz/go-full/pkg/deebee"
)

// SqlPublisherAdapter provides an adapter to publish messages to a SQL database.
type SqlPublisherAdapter struct {
	*sql.Publisher
}

// NewSqlPublisherAdapter creates a new SqlPublisherAdapter using the provided database connection.
func NewSqlPublisherAdapter(db deebee.IDatabase) (SqlPublisherAdapter, error) {
	publisher, err := sql.NewPublisher(
		sql.BeginnerFromPgx(db),
		sql.PublisherConfig{
			SchemaAdapter: sql.DefaultPostgreSQLSchema{},
		},
		nil,
	)
	if err != nil {
		return SqlPublisherAdapter{}, err
	}

	return SqlPublisherAdapter{
		Publisher: publisher,
	}, nil
}
