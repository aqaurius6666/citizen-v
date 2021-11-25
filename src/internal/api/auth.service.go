package api

import (
	"time"

	"github.com/aquarius6666/citizen-v/src/internal/db"
	"github.com/aquarius6666/citizen-v/src/internal/db/role"
	"github.com/aquarius6666/citizen-v/src/internal/db/user"
	"github.com/aquarius6666/citizen-v/src/internal/services/jwt"
	"github.com/aquarius6666/citizen-v/src/internal/var/c"
	"github.com/aquarius6666/citizen-v/src/internal/var/e"
	"github.com/aquarius6666/citizen-v/src/pb"
	"golang.org/x/xerrors"
)

type AuthService struct {
	Repo       db.ServerRepo
	JWTService jwt.JWT
}

func (s *AuthService) Register(req *pb.PostRegisterRequest) (*pb.PostRegisterResponse_Data, error) {
	var err error
	var u *user.User
	var r *role.Role
	if req.Username == "" || req.Password == "" {
		return nil, e.ErrMissingBody
	}
	if _, err = s.Repo.SelectUser(&user.Search{
		User: user.User{
			Username: &req.Username,
		},
	}); err == nil {
		return nil, e.ErrAuthUsernameExisted
	}
	if r, err = s.Repo.SelectRole(&role.Search{
		Role: role.Role{
			Name: &role.ROLE_A1,
		},
	}); err != nil {
		return nil, err
	}
	if u, err = s.Repo.InsertUser(&user.User{
		Username:     &req.Username,
		HashPassword: &req.Password,
		RoleID:       r.ID,
	}); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.PostRegisterResponse_Data{
		Username: *u.Username,
		Password: req.Password,
		Id:       u.ID.String(),
	}, nil
}

func (s *AuthService) Login(req *pb.PostLoginRequest) (*pb.PostLoginResponse_Data, error) {
	var err error
	var u *user.User
	if req.Username == "" || req.Password == "" {
		return nil, e.ErrMissingBody
	}
	if u, err = s.Repo.SelectUser(&user.Search{
		User: user.User{
			Username: &req.Username,
		},
	}); err != nil {
		return nil, err
	}
	if *u.HashPassword != req.Password {
		return nil, e.ErrAuthCredentialWrong
	}
	token, err := s.JWTService.Sign(jwt.ClaimStruct{
		ExpiresAt: time.Now().Add(c.JWT_EXPIRED_DURATION),
		Uid:       u.ID.String(),
		IssuedAt:  time.Now(),
		Issuer:    c.JWT_ISSUER,
		NotBefore: time.Now(),
		Subject:   c.JWT_SUBJECT,
		RoleName:  *u.Role.Name,
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.PostLoginResponse_Data{
		Token: token,
	}, nil
}
