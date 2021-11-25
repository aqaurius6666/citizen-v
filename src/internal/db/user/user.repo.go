package user

type RoleRepo interface {
	SelectUser(*Search) (*User, error)
	InsertUser(*User) (*User, error)
}
