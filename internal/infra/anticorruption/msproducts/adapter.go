package msproducts

import (
	"context"
	"encoding/json"
	"github.com/braiphub/go-core/queue"
	"github.com/braiphub/ms-tech-talk/internal/domain/entity"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type OfferService interface {
	Upsert(ctx context.Context, offer entity.Offer) error
}

type Adapter struct {
	queue        queue.QueueI
	offerService OfferService
	validator    *validator.Validate
}

func NewAdapter(queue queue.QueueI, offerService OfferService) *Adapter {
	return &Adapter{
		queue:        queue,
		offerService: offerService,
		validator:    validator.New(),
	}
}

func (a *Adapter) StartConsumers(ctx context.Context) {
	go a.queue.Consume(ctx, "ms-tech-talk.ms-products.offer-created", a.HandleOrderCreatedEvent)
}

func (a *Adapter) HandleOrderCreatedEvent(ctx context.Context, msg queue.Message) error {
	var request CreateOfferRequestDTO

	if err := json.Unmarshal(msg.Body, &request); err != nil {
		return errors.Wrap(err, "unmarshal")
	}

	if err := a.validator.Struct(request); err != nil {
		return errors.Wrap(err, "validate")
	}

	offer := translateCreateOfferRequestDTOToEntity(request)

	if err := a.offerService.Upsert(ctx, offer); err != nil {
		return errors.Wrap(err, "service")
	}

	return nil
}
