package user

type Repository interface {
	Get() ([]*User, error)
}

type RepositoryImpl struct {}

func (r *RepositoryImpl) Get() ([]*User, error) {
	return nil, nil
}
