package jwt

import "time"

type SecretKey string

type JWT interface {
	Sign(ClaimStruct) (string, error)
	Verify(string) (bool, map[string]interface{}, error)
}

func NewJWT(sec SecretKey) JWT {
	return NewJWTService(sec)
}

type ClaimStruct struct {
	ExpiresAt time.Time
	Uid       string
	IssuedAt  time.Time
	RoleName  string
}
