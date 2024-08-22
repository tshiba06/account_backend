package role

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Get(ctx context.Context) ([]*MasterRole, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) Get(ctx context.Context) ([]*MasterRole, error) {
	var roles []*MasterRole
	if err := r.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
