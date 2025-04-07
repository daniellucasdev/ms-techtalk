package repository

import (
	"context"
	"github.com/braiphub/ms-tech-talk/internal/domain/entity"
	"gorm.io/gorm"
)

type ReadOfferRepository struct {
	db *gorm.DB
}

func NewReadOfferRepository(gorm *gorm.DB) *ReadOfferRepository {
	return &ReadOfferRepository{
		db: gorm,
	}
}

func (r *ReadOfferRepository) FindByHash(ctx context.Context, offerHash string) (*entity.Offer, error) {
	var offer entity.Offer
	if err := r.db.Model(&offer).
		Where("hash = ?", offer.Hash).
		First(&offer).Error; err != nil {
	}

	return &offer, nil
}
