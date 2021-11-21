package user

type UserRepo interface {
	SelectUser(*Search) (*User, error)
	InsertUser(*User) (*User, error)
}
