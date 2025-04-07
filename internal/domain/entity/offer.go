package entity

import (
	"github.com/braiphub/ms-tech-talk/internal/domain/enum"
	"gorm.io/gorm"
)

type Offer struct {
	gorm.Model
	Hash        string
	Name        string
	ProductType enum.ProductType
}

func NewOffer(hash, name string, productType enum.ProductType) Offer {
	return Offer{
		Hash:        hash,
		Name:        name,
		ProductType: productType,
	}
}
