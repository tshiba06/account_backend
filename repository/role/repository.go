package role

type Repository interface {
	Get() ([]*Role, error)
}

type RepositoryImpl struct {}

func (r *RepositoryImpl) Get() ([]*Role, error) {
	return nil, nil
}
