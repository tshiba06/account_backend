package user

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type Repository interface {
	Get(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, entity *User) error
	Update(ctx context.Context, entity *User) error
}

type RepositoryImpl struct {
	db     *gorm.DB
	tracer trace.Tracer
}

func NewRepository(db *gorm.DB, tracer trace.Tracer) Repository {
	return &RepositoryImpl{
		db:     db,
		tracer: tracer,
	}
}

func (r *RepositoryImpl) Get(ctx context.Context) ([]*User, error) {
	_, span := r.tracer.Start(ctx, "UserGetRepository")
	defer span.End()

	return nil, nil
}

func (r *RepositoryImpl) Create(ctx context.Context, entity *User) error {
	_, span := r.tracer.Start(ctx, "UserCreateRepository")
	defer span.End()

	return nil
}

func (r *RepositoryImpl) Update(ctx context.Context, entity *User) error {
	_, span := r.tracer.Start(ctx, "UserUpdateRepository")
	defer span.End()

	return nil
}
