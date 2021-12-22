package api

import (
	"time"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/services/jwt"
	"github.com/aqaurius6666/citizen-v/src/internal/var/c"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type AuthService struct {
	Repo       db.ServerRepo
	JWTService jwt.JWT
	Model      model.Server
	Logger     *logrus.Logger
}

func (s *AuthService) Register(req *pb.PostRegisterRequest) (*pb.PostRegisterResponse_Data, error) {
	var err error
	var u *user.User
	var r *role.Role
	if f, ok := validate.RequiredFields(req, "Password", "Username"); !ok {
		return nil, e.ErrMissingField(f)
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

func (s *AuthService) Auth(req *pb.GetAuthRequest) (*pb.GetAuthResponse_Data, error) {
	var err error
	var u *user.User
	if f, ok := validate.RequiredFields(req, "XCallerId"); !ok {
		return nil, e.ErrMissingField(f)
	}
	if u, err = s.Repo.SelectUser(&user.Search{
		User: user.User{
			BaseModel: database.BaseModel{ID: uuid.MustParse(req.XCallerId)},
		},
	}); err != nil {
		return nil, err
	}
	return &pb.GetAuthResponse_Data{
		User: lib.ConvertOneUser(u, s.Repo),
	}, nil
}

func (s *AuthService) Login(req *pb.PostLoginRequest) (*pb.PostLoginResponse_Data, error) {
	var err error
	var u *user.User
	if f, ok := validate.RequiredFields(req, "Username", "Password"); !ok {
		return nil, e.ErrMissingField(f)
	}
	if u, err = s.Repo.SelectUser(&user.Search{
		User: user.User{
			Username:     &req.Username,
			HashPassword: lib.MyHashPassword(&req.Password),
		},
	}); err != nil {
		return nil, e.ErrAuthCredentialWrong
	}
	token, err := s.JWTService.Sign(jwt.ClaimStruct{
		ExpiresAt: time.Now().Add(c.JWT_EXPIRED_DURATION),
		Uid:       u.ID.String(),
		IssuedAt:  time.Now(),
		RoleName:  *u.Role.Name,
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.PostLoginResponse_Data{
		Token:              token,
		UseDefaultPassword: *u.UseDefaultPassword,
	}, nil
}

func (s *AuthService) ChangePassword(req *pb.PostAuthPasswordRequest) (*pb.PostAuthPasswordResponse_Data, error) {
	var err error
	var u *user.User
	if f, ok := validate.RequiredFields(req, "OldPassword", "NewPassword", "XCallerId"); !ok {
		return nil, e.ErrMissingField(f)
	}
	if u, err = s.Repo.SelectUser(&user.Search{
		User: user.User{
			BaseModel:    database.BaseModel{ID: uuid.MustParse(req.XCallerId)},
			HashPassword: lib.MyHashPassword(&req.OldPassword),
		},
	}); err != nil {
		return nil, e.ErrAuthCredentialWrong
	}
	if err = s.Repo.UpdateUser(&user.Search{
		User: user.User{
			BaseModel: u.BaseModel,
		},
	}, &user.User{
		HashPassword:       lib.MyHashPassword(&req.NewPassword),
		UseDefaultPassword: utils.BoolPtr(false),
	}); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.PostAuthPasswordResponse_Data{}, nil
}
