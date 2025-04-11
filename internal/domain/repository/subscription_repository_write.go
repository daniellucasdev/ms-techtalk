package repository

import (
	"context"
	"github.com/braiphub/go-core/hashid"
	"github.com/braiphub/ms-tech-talk/internal/domain/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type WriteSubscriptionRepository struct {
	db     *gorm.DB
	hasher hashid.Hasher
}

func NewWriteSubscriptionRepository(db *gorm.DB, hasher hashid.Hasher) *WriteSubscriptionRepository {
	return &WriteSubscriptionRepository{
		db:     db,
		hasher: hasher.WithPrefix("sub"),
	}
}

func (s *WriteSubscriptionRepository) Create(ctx context.Context, subscription *entity.Subscription) error {
	if err := s.db.WithContext(ctx).Create(subscription).Error; err != nil {
		return errors.Wrap(err, "db create")
	}

	hash, err := s.hasher.Generate(subscription.ID)
	if err != nil {
		return errors.Wrap(err, "hasher")
	}

	subscription.Hash = hash
	if err := s.db.WithContext(ctx).Save(subscription).Error; err != nil {
		return errors.Wrap(err, "db save")
	}

	return nil
}
