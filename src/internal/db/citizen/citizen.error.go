package citizen

import "golang.org/x/xerrors"

var (
	PREFIX        = "citizen"
	ErrNotFound   = xerrors.Errorf("%s: record not found", PREFIX)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", PREFIX)
)
