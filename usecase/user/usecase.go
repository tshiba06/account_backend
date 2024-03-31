package user

import (
	userRepository "github.com/tshiba06/account_backend/repository/user"
)

type UseCase interface {
	Create(params User) error
	Update(params User) error
	Get() ([]*User, error)
}

type UseCaseImpl struct {
	userRepo userRepository.Repository
}

func NewUsecase(userRepo userRepository.Repository) UseCase {
	return &UseCaseImpl{
		userRepo: userRepo,
	}
}

func (u *UseCaseImpl) Create(params User) error {
	return nil
}

func (u *UseCaseImpl) Update(params User) error {
	return nil
}

func (u *UseCaseImpl) Get() ([]*User, error) {
	return nil, nil
}
