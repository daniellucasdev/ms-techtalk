package event

const (
	SubscriptionCreated = "subscription.created"
)

type SubscriptionCreatedEvent struct {
	Hash      string
	OfferHash string
}
