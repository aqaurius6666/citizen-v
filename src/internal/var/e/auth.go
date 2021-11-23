package e

import "golang.org/x/xerrors"

var (
	ErrAuthUsernameExisted = xerrors.New("username has already existed")
	ErrAuthCredentialWrong = xerrors.New("username or password wrong")
	ErrAuthTokenFail       = xerrors.New("token fail")
)
