package e

import "golang.org/x/xerrors"

var (
	ErrAuthUsernameExisted      = xerrors.New("username has already existed")
	ErrAuthCredentialWrong      = xerrors.New("username or password wrong")
	ErrAuthTokenFail            = xerrors.New("token fail")
	ErrAuthMissingAuthorization = xerrors.New("missing authorization token")
	ErrAuthTokenInvalid         = xerrors.New("Token is invalid")
)
