package e

import (
	"golang.org/x/xerrors"
)

var (
	ErrBodyInvalid        = xerrors.New("body invalid")
	ErrMissingBody        = xerrors.New("missing body")
	ErrIdInvalid          = xerrors.New("id invalid")
	ErrZoneAccountExisted = xerrors.New("zone account existed")
)
