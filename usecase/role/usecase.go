package role

import (
	"context"

	roleRepository "github.com/tshiba06/account_backend/repository/role"
)

type UseCase interface {
	Get(ctx context.Context) ([]*Role, error)
}

type UseCaseImpl struct {
	roleRepo roleRepository.Repository
}

func NewUseCase(roleRepo roleRepository.Repository) UseCase {
	return &UseCaseImpl{
		roleRepo: roleRepo,
	}
}

func (u UseCaseImpl) Get(ctx context.Context) ([]*Role, error) {
	roles, err := u.roleRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	var results []*Role
	for _, r := range roles {
		results = append(results, &Role{
			ID:          r.ID,
			Name:        r.Name,
			DisplayName: r.DisplayName,
		})
	}

	return results, nil
}
