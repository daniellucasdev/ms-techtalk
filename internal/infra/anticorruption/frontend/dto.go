package frontend

type CreateSubscriptionRequestDTO struct {
	OrderHash string `json:"order_hash" validate:"required"`
	OfferHash string `json:"offer_hash" validate:"required"`
	Amount    int    `json:"amount"     validate:"required"`
}

type CreateSubscriptionResponseDTO struct {
	Hash      string `json:"hash"`
	OrderHash string `json:"order_hash"`
	OfferHash string `json:"offer_hash"`
	Amount    int    `json:"amount"`
}
