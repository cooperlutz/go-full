package inbound

import (
	"github.com/cooperlutz/go-full/internal/retailsales/app"
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
		"retailsales_item_added_to_cart_handler",
		"retailsales.item_added_to_cart",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"retailsales_item_removed_from_cart_handler",
		"retailsales.item_removed_from_cart",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"retailsales_sales_order_placed_handler",
		"retailsales.sales_order_placed",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"retailsales_inventory_deduction_requested_handler",
		"retailsales.inventory_deduction_requested",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"retailsales_sales_order_fulfilled_handler",
		"retailsales.sales_order_fulfilled",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"retailsales_sales_order_canceled_handler",
		"retailsales.sales_order_canceled",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"retailsales_inventory_restock_requested_handler",
		"retailsales.inventory_restock_requested",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)
}
