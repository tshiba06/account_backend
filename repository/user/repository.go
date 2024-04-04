package user

import "github.com/jmoiron/sqlx"

type Repository interface {
	Get() ([]*User, error)
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) Get() ([]*User, error) {

	return nil, nil
}
