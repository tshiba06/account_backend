package role

type Repository interface {
	Get() ([]*Role, error)
}
