package user

type UserRepo interface {
	SelectOne(*Search) (*User, error)
}
