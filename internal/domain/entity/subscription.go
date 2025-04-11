package entity

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	Hash      string
	OfferID   uint
	Offer     Offer
	OrderHash string
	Amount    int
}

func NewSubscription(orderHash string, amount int, offer Offer) *Subscription {
	return &Subscription{
		OrderHash: orderHash,
		Offer:     offer,
		OfferID:   offer.ID,
		Amount:    amount,
	}
}
