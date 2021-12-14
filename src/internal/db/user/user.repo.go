package user

type UserRepo interface {
	SelectUser(*Search) (*User, error)
	InsertUser(*User) (*User, error)
	CountUser(*Search) (*int64, error)
	UpdateUser(*Search, *User) error
}
