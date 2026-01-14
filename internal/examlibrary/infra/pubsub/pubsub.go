package pubsub

import (
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/repository"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

// Ensure EventProcessor implements the IPubSubEventProcessor interface.
var _ eeventdriven.IPubSubEventProcessor = (*examLibraryPubSub)(nil)

// examLibraryPubSub handles Pub/Sub events specific to the ExamLibrary module.
type examLibraryPubSub struct {
	*eeventdriven.BasePgsqlPubSubProcessor
	repo repository.IExamLibraryRepository
}

// New - Creates a new instance of examLibraryPubSub with the provided database connection and repository.
func New(db deebee.IDatabase, repo repository.IExamLibraryRepository) (*examLibraryPubSub, error) {
	basePS, err := eeventdriven.NewPubSub(db)
	if err != nil {
		return nil, err
	}

	ps := &examLibraryPubSub{
		BasePgsqlPubSubProcessor: basePS,
		repo:                     repo,
	}

	return ps, nil
}

// RegisterHandlers overrides the base method to register ExamLibrary-specific event handlers.
func (pp *examLibraryPubSub) RegisterSubscriberHandlers() error {
	router := pp.GetRouter()

	router.AddConsumerHandler(
		"handler_exam_library_new_exam_created",
		"examlibrary",
		pp.GetSubscriber(),
		eeventdriven.TraceConsumerHandler(func(msg *message.Message) error {
			msg.Ack()

			return nil
		}),
	)

	return nil
}
