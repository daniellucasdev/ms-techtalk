package handlers

import (
	"github.com/asaskevich/EventBus"
	"github.com/braiphub/go-core/log"
	"github.com/braiphub/ms-tech-talk/internal/domain/service"
	"github.com/braiphub/ms-tech-talk/internal/events/bus"
	"github.com/braiphub/ms-tech-talk/internal/events/event"
)

type EventHandler struct {
	logger              log.LoggerI
	bus                 EventBus.Bus
	offerService        service.OfferService
	subscriptionService service.SubscriptionService
}

func NewEventHandler(
	logger log.LoggerI,

) *EventHandler {
	return &EventHandler{
		logger: logger,
		bus:    bus.GetBus(),
	}
}

//nolint:errcheck // it'll be changed soon with another implementation with error logging
func (handler *EventHandler) StartListeners() {
	handler.bus.SubscribeAsync(event.OfferCreated, handler.LogOfferCreated, false)

	handler.bus.SubscribeAsync(event.SubscriptionCreated, handler.LogSubscriptionCreated, false)
}
