package user

type User struct {
	ID int
	Name string
	Email string
	MasterRoleID int
	CreatedAt time.Time
	UpdatedAt time.Time
}
