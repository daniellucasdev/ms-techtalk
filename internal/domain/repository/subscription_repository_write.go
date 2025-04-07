package repository

import (
	"context"
	"github.com/braiphub/ms-tech-talk/internal/domain/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type WriteSubscriptionRepository struct {
	db *gorm.DB
}

func NewWriteSubscriptionRepository(db *gorm.DB) *WriteSubscriptionRepository {
	return &WriteSubscriptionRepository{
		db: db,
	}
}

func (s *WriteSubscriptionRepository) Create(ctx context.Context, subscription *entity.Subscription) error {
	if err := s.db.WithContext(ctx).Create(subscription).Error; err != nil {
		return errors.Wrap(err, "repo")
	}

	return nil
}
