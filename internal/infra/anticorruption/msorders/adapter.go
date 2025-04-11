package msorders

import (
    "context"
    "database/sql"
    "encoding/json"
    "github.com/braiphub/go-core/log"
    "github.com/braiphub/go-core/queue"
    "github.com/braiphub/ms-tech-talk/internal/domain/entity"
    "github.com/braiphub/ms-tech-talk/internal/events/bus"
    "github.com/pkg/errors"
)

type OfferService interface {
    FindByHash(ctx context.Context, hash string) (*entity.Offer, error)
}

type Adapter struct {
    queue    queue.QueueI
    offerSvc OfferService
    bus      bus.EventBusI
    logger   log.LoggerI
}

func NewAdapter(
    queue queue.QueueI,
    offerScv OfferService,
    logger log.LoggerI,
) *Adapter {
    return &Adapter{
        queue:    queue,
        offerSvc: offerScv,
        bus:      bus.GetBus(),
        logger:   logger,
    }
}

func (a *Adapter) StartConsumers(ctx context.Context) {
    go a.queue.Consume(ctx, "ms-tech-talk.ms-orders.order-created", a.HandleOrderCreated)
}

type OrderCreatedDTO struct {
    OrderHash string `json:"order_hash"`
    OfferHash string `json:"offer_hash"`
    Ammount   int    `json:"ammount"`
}

func (a *Adapter) HandleOrderCreated(ctx context.Context, message queue.Message) error {
    var dto OrderCreatedDTO
    if err := json.Unmarshal(message.Body, &dto); err != nil {
        return err
    }

    offer, err := a.offerSvc.FindByHash(ctx, dto.OfferHash)

    if errors.Is(err, sql.ErrNoRows) {
        return nil
    }

    if err != nil {
        return err
    }

    a.bus.Publish("order-created", offer)
    a.logger.Info("Order created !!")

    return nil
}