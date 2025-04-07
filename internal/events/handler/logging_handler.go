package handlers

import (
	"github.com/braiphub/go-core/log"
	"github.com/braiphub/ms-tech-talk/internal/events/event"
)

func (handler *EventHandler) LogOfferCreated(ev event.OfferCreatedEvent) {
	handler.logger.Info("offer created: ", log.Any("offer_hash", ev.Hash))
}

func (handler *EventHandler) LogSubscriptionCreated(ev event.SubscriptionCreatedEvent) {
	handler.logger.Info("subscription created: ", log.Any("subscription_hash", ev.Hash))
}
