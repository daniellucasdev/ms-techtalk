package repository

import (
	"context"
	"database/sql"
	"github.com/braiphub/ms-tech-talk/internal/domain/entity"
	"github.com/pkg/errors"
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
	err := r.db.Model(&offer).
		Where("hash = ?", offerHash).
		First(&offer).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, errors.Wrap(err, "db")
	}

	return &offer, nil
}
