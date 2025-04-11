package service

import (
	"context"
	"github.com/braiphub/ms-tech-talk/internal/domain/entity"
	"github.com/braiphub/ms-tech-talk/internal/events/bus"
	"github.com/braiphub/ms-tech-talk/internal/events/event"
	"github.com/pkg/errors"
)

type SubscriptionWriteRepository interface {
	Create(ctx context.Context, subscription *entity.Subscription) error
}

type SubscriptionService struct {
	writeRepo SubscriptionWriteRepository
	bus       bus.EventBusI
}

func NewSubscriptionService(writeRepo SubscriptionWriteRepository) *SubscriptionService {
	return &SubscriptionService{
		writeRepo: writeRepo,
		bus:       bus.GetBus(),
	}
}

func (s *SubscriptionService) Create(
	ctx context.Context,
	subscription *entity.Subscription,
) error {
	if err := s.writeRepo.Create(ctx, subscription); err != nil {
		return errors.Wrap(err, "repo")
	}

	s.bus.Publish(event.SubscriptionCreated, event.SubscriptionCreatedEvent{
		Hash:      subscription.Hash,
		OfferHash: subscription.Offer.Hash,
	})

	return nil
}
