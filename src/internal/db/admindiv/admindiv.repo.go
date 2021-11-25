package admindiv

type AdminDivRepo interface {
	SelectAdminDiv(*Search) (*AdminDiv, error)
	InsertAdminDiv(*AdminDiv) (*AdminDiv, error)
}
