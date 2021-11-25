package role

import "golang.org/x/xerrors"

var (
	PREFIX        = "role"
	ErrNotFound   = xerrors.Errorf("%s: record not found", PREFIX)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", PREFIX)
)
