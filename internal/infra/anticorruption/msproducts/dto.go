package msproducts

import "github.com/braiphub/ms-tech-talk/internal/domain/enum"

type CreateOfferRequestDTO struct {
	OfferHash   string           `json:"offer_hash"   validate:"required"`
	OfferName   string           `json:"offer_name"   validate:"required"`
	ProductType enum.ProductType `json:"product_type" validate:"required,oneof=PHYSICAL DIGITAL"`
}
