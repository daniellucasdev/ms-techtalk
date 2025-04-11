package msorders

import "github.com/braiphub/go-core/queue"

type OfferService interface {
	Upsert(ctx context.Context, hash string) (*entity.Offer, error)	
}

type Adapter struct {
	queue queue.QueueI
	offerSvc OfferService
}

func NewAdapter (queue queue.QueueI, offerSvc OfferService) *Adapter {
	return &Adapter{
		queue: queue,
		offerSvc: offerSvc,
	}
}