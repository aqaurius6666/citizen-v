package api

import (
	"fmt"
	"time"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/services/jwt"
	"github.com/aqaurius6666/citizen-v/src/internal/var/c"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type AuthService struct {
	Repo       db.ServerRepo
	JWTService jwt.JWT
}

func (s *AuthService) Issue(req *pb.PostAuthIssueRequest) (*pb.PostAuthIssueResponse_Data, error) {
	var err error
	if !validate.RequiredFields(req, "AdminDivId", "RoleId") {
		return nil, e.ErrMissingBody
	}
	var aid, rid uuid.UUID
	if aid, err = uuid.Parse(req.AdminDivId); err != nil {
		return nil, e.ErrIdInvalid
	}
	if rid, err = uuid.Parse(req.RoleId); err != nil {
		return nil, e.ErrIdInvalid
	}
	add, err := s.Repo.SelectAdminDiv(&admindiv.Search{
		AdminDiv: admindiv.AdminDiv{
			BaseModel: database.BaseModel{ID: aid},
		},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	rol, err := s.Repo.SelectRole(&role.Search{
		Role: role.Role{
			BaseModel: database.BaseModel{ID: rid},
		},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	number, err := s.Repo.CountUser(&user.Search{
		User: user.User{
			AdminDivID: aid,
		},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	fmt.Printf("number: %v\n", *number)
	if *number > 0 {
		return nil, e.ErrZoneAccountExisted
	}
	pass := lib.RandomPassword()
	tmp := user.User{
		Username:     utils.StrPtr(fmt.Sprintf("citizen%s", *add.Code)),
		HashPassword: lib.MyHashPassword(pass),
		AdminDivID:   add.ID,
		RoleID:       rol.ID,
	}
	usr, err := s.Repo.InsertUser(&tmp)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.PostAuthIssueResponse_Data{
		Username: *usr.Username,
		Password: *pass,
	}, nil
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
	if !validate.RequiredFields(req, "Username", "Password") {
		return nil, e.ErrMissingBody
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
	if !validate.RequiredFields(req, "OldPassword", "NewPassword", "Id") {
		return nil, e.ErrMissingBody
	}
	if u, err = s.Repo.SelectUser(&user.Search{
		User: user.User{
			BaseModel:    database.BaseModel{ID: uuid.MustParse(req.Id)},
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
