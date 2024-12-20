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

	var users []*User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *RepositoryImpl) Create(ctx context.Context, entity *User) error {
	_, span := r.tracer.Start(ctx, "UserCreateRepository")
	defer span.End()

	if err := r.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return err
	}

	return nil
}

func (r *RepositoryImpl) Update(ctx context.Context, entity *User) error {
	_, span := r.tracer.Start(ctx, "UserUpdateRepository")
	defer span.End()

	if err := r.db.WithContext(ctx).Updates(&entity).Error; err != nil {
		return err
	}

	return nil
}
