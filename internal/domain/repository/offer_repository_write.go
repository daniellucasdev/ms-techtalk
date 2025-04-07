package repository

import (
	"context"
	"github.com/braiphub/ms-tech-talk/internal/domain/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type WriteOfferRepository struct {
	db *gorm.DB
}

func NewWriteOfferRepository(db *gorm.DB) *WriteOfferRepository {
	return &WriteOfferRepository{
		db: db,
	}
}

func (w *WriteOfferRepository) Save(ctx context.Context, offer *entity.Offer) error {
	err := w.db.WithContext(ctx).Save(offer).Error
	if err != nil {
		return errors.Wrap(err, "repo")
	}

	return nil
}
