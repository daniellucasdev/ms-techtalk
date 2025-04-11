package frontend

import (
	"github.com/braiphub/ms-tech-talk/internal/domain/entity"
)

func TranslateCreateSubscriptionRequestToEntity(dto CreateSubscriptionRequestDTO) entity.Subscription {
	return entity.Subscription{
		OrderHash: dto.OrderHash,
		Amount:    dto.Amount,
	}
}

func TranslateSubscriptionToCreateSubscriptionResponse(subscription *entity.Subscription) entity.Subscription {
	return entity.Subscription{
		Hash:      subscription.Hash,
		OrderHash: subscription.OrderHash,
		Amount:    subscription.Amount,
	}
}
