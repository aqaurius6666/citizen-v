package campaign

import "golang.org/x/xerrors"

var (
	PREFIX        = "campaign"
	ErrNotFound   = xerrors.Errorf("%s: record not found", PREFIX)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", PREFIX)
	ErrUpdateFail = xerrors.Errorf("%s: update failed", PREFIX)
	ErrInvalid    = xerrors.Errorf("%s: invalid", PREFIX)
)
