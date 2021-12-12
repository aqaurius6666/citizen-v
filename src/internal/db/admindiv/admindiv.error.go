package admindiv

import "golang.org/x/xerrors"

var (
	PREFIX        = "administrative_division"
	ErrNotFound   = xerrors.Errorf("%s: record not found", PREFIX)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", PREFIX)
	ErrInvalid    = xerrors.Errorf("%s: invalid", PREFIX)
)
