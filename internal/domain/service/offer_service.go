package service

import (
	"context"
	"github.com/braiphub/ms-tech-talk/internal/domain/entity"
	"github.com/braiphub/ms-tech-talk/internal/events/bus"
	"github.com/braiphub/ms-tech-talk/internal/events/event"
	"github.com/pkg/errors"
)

type OfferWriteRepository interface {
	Save(ctx context.Context, offer *entity.Offer) error
}

type OfferReadRepository interface {
	FindByHash(ctx context.Context, offerHash string) (*entity.Offer, error)
}
type OfferService struct {
	repoWrite OfferWriteRepository
	repoRead  OfferReadRepository
	bus       bus.EventBusI
}

func NewOfferService(repoWrite OfferWriteRepository, repoRead OfferReadRepository) *OfferService {
	return &OfferService{
		repoWrite: repoWrite,
		repoRead:  repoRead,
		bus:       bus.GetBus(),
	}
}

func (s *OfferService) Upsert(ctx context.Context, offer entity.Offer) error {
	if err := s.repoWrite.Save(ctx, &offer); err != nil {
		return errors.Wrap(err, "repo")
	}

	s.bus.Publish(event.OfferCreated, event.OfferCreatedEvent{
		Hash: offer.Hash,
	})

	return nil
}

func (s *OfferService) FindByHash(ctx context.Context, offerHash string) (*entity.Offer, error) {
	offer, err := s.repoRead.FindByHash(ctx, offerHash)
	if err != nil {
		return nil, errors.Wrap(err, "repo")
	}

	return offer, nil
}
