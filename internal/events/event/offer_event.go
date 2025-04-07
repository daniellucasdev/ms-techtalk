package event

const (
	OfferCreated = "offer.created"
)

type OfferCreatedEvent struct {
	Hash string
}
