package msproducts

import "github.com/braiphub/ms-tech-talk/internal/domain/entity"

func translateCreateOfferRequestDTOToEntity(dto CreateOfferRequestDTO) entity.Offer {
	return entity.NewOffer(dto.OfferHash, dto.OfferName, dto.ProductType)
}
