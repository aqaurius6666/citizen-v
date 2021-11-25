package user

import "golang.org/x/xerrors"

var (
	PREFIX             = "user"
	ErrNotFound        = xerrors.Errorf("%s: record not found", PREFIX)
	ErrInsertFail      = xerrors.Errorf("%s: insert failed", PREFIX)
	ErrUsernameExisted = xerrors.Errorf("%s: username has already existed", PREFIX)
)
