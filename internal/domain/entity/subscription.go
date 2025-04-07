package entity

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	Hash      string
	OrderHash string
	OfferHash string
	Amount    int
}
