package user

import "golang.org/x/xerrors"

var (
	ErrNotFound        = xerrors.New("record not found")
	ErrInsertFail      = xerrors.New("insert failed")
	ErrUsernameExisted = xerrors.New("username has already existed")
)
