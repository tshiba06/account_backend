package user

type Repository interface {
	Get() ([]*User, error)
}
