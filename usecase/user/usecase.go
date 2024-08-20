package user

import (
	"context"

	userRepository "github.com/tshiba06/account_backend/repository/user"
	"go.opentelemetry.io/otel/trace"
)

type UseCase interface {
	Create(ctx context.Context, params User) error
	Update(ctx context.Context, params User) error
	Get(ctx context.Context) ([]*User, error)
}

type UseCaseImpl struct {
	tracer   trace.Tracer
	userRepo userRepository.Repository
}

func NewUsecase(
	tracer trace.Tracer,
	userRepo userRepository.Repository,
) UseCase {
	return &UseCaseImpl{
		tracer:   tracer,
		userRepo: userRepo,
	}
}

func (u *UseCaseImpl) Create(ctx context.Context, params User) error {
	ctx, span := u.tracer.Start(ctx, "UseCreateUseCase")
	defer span.End()

	return nil
}

func (u *UseCaseImpl) Update(ctx context.Context, params User) error {
	return nil
}

func (u *UseCaseImpl) Get(ctx context.Context) ([]*User, error) {
	ctx, span := u.tracer.Start(ctx, "UserGetUseCase")
	defer span.End()

	// TODO: 中身実装後にusersを返す
	_, err := u.userRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
