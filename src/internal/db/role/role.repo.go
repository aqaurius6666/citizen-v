package role

type RoleRepo interface {
	SelectRole(*Search) (*Role, error)
	InsertRole(*Role) (*Role, error)
}
