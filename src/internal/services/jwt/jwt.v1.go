package jwt

import (
	"fmt"
	"strings"

	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/xerrors"
)

var (
	_ JWT = JWTService{}
)

type JWTService struct {
	secretKey string
}

func NewJWTService(sec SecretKey) JWTService {
	return JWTService{
		secretKey: string(sec),
	}
}

func (s JWTService) Sign(c ClaimStruct) (string, error) {
	claim := jwt.MapClaims{
		"exp":  jwt.NewNumericDate(c.ExpiresAt),
		"iat":  jwt.NewNumericDate(c.IssuedAt),
		"uid":  c.Uid,
		"role": c.RoleName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}

	return tokenString, nil
}
func (s JWTService) Verify(tokenString string) (bool, map[string]interface{}, error) {
	var idString string

	if len(strings.Split(tokenString, ".")) != 3 {
		return false, nil, e.ErrAuthTokenInvalid
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return false, nil, xerrors.Errorf("%w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, nil, xerrors.Errorf("%w", e.ErrAuthTokenInvalid)
	}
	if err := claims.Valid(); err != nil {
		return false, nil, xerrors.Errorf("%w", err)
	}
	if idString, ok = claims["uid"].(string); idString == "" || !ok {
		return false, nil, xerrors.Errorf("%w", e.ErrAuthTokenInvalid)
	}
	return true, claims, nil
}
