package jwt

type SecretKey string

type JWTService struct {
	secretKey string
}

func NewJWTService(sec SecretKey) JWTService {
	return JWTService{
		secretKey: string(sec),
	}
}

func (s JWTService) Sign() (string, error) {

	return "lewlew", nil
}

func (s JWTService) Verify() (bool, error) {

	return true, nil
}
