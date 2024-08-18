package role

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Get() ([]*MasterRole, error)
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) Get() ([]*MasterRole, error) {
	return nil, nil
}
