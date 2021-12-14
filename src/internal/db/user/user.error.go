package user

import "golang.org/x/xerrors"

var (
	PREFIX             = "user"
	ErrNotFound        = xerrors.Errorf("%s: record not found", PREFIX)
	ErrInsertFail      = xerrors.Errorf("%s: insert failed", PREFIX)
	ErrUsernameExisted = xerrors.Errorf("%s: username has already existed", PREFIX)
	ErrUpdateFail      = xerrors.Errorf("%s: update failed", PREFIX)
	ErrInvalid         = xerrors.Errorf("%s: invalid", PREFIX)
)
