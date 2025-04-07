package frontend

import (
	"github.com/braiphub/ms-tech-talk/internal/domain/entity"
)

func TranslateCreateSubscriptionRequestToEntity(dto CreateSubscriptionRequestDTO) entity.Subscription {
	return entity.Subscription{
		OrderHash: dto.OrderHash,
		OfferHash: dto.OfferHash,
		Amount:    dto.Amount,
	}
}

func TranslateSubscriptionToCreateSubscriptionResponse(subscription *entity.Subscription) entity.Subscription {
	return entity.Subscription{
		Hash:      subscription.Hash,
		OrderHash: subscription.OrderHash,
		OfferHash: subscription.OfferHash,
		Amount:    subscription.Amount,
	}
}
