package inbound

import (
	"github.com/cooperlutz/go-full/internal/insuranceandclaims/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type SqlSubscriberAdapter struct {
	app    app.Events
	pubSub *eeventdriven.BasePgsqlPubSubProcessor
}

func NewSqlSubscriberAdapter(events app.Events, pubSub *eeventdriven.BasePgsqlPubSubProcessor) SqlSubscriberAdapter {
	return SqlSubscriberAdapter{
		app:    events,
		pubSub: pubSub,
	}
}

func (s SqlSubscriberAdapter) RegisterEventHandlers() {
	router := s.pubSub.GetRouter()

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"insuranceandclaims_insurance_provider_registered_handler",
		"insuranceandclaims.insurance_provider_registered",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"insuranceandclaims_insurance_claim_submitted_handler",
		"insuranceandclaims.insurance_claim_submitted",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"insuranceandclaims_insurance_claim_approved_handler",
		"insuranceandclaims.insurance_claim_approved",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"insuranceandclaims_insurance_claim_rejected_handler",
		"insuranceandclaims.insurance_claim_rejected",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)
}
