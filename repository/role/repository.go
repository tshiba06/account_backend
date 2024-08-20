package role

import (
	"gorm.io/gorm"
)

type Repository interface {
	Get() ([]*MasterRole, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) Get() ([]*MasterRole, error) {
	return nil, nil
}
