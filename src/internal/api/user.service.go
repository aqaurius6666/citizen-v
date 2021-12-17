package api

import (
	"fmt"
	"strconv"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aqaurius6666/citizen-v/src/internal/db/role"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/lib/validate"
	"github.com/aqaurius6666/citizen-v/src/internal/model"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/aqaurius6666/go-utils/database"
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type UserService struct {
	Repo  db.ServerRepo
	Model model.Server
}

func (s *UserService) Active(req *pb.PostUserActiveRequest) (*pb.PostUserActiveResponse_Data, error) {
	if f, ok := validate.RequiredFields(req, "Id", "Value", "CallerId"); !ok {
		return nil, e.ErrMissingField(f)
	}
	var err error
	var uid uuid.UUID
	if uid, err = uuid.Parse(req.Id); err != nil {
		return nil, e.ErrIdInvalid
	}
	search := &user.Search{
		User: user.User{
			BaseModel: database.BaseModel{ID: uid},
		},
	}
	usr, err := s.Repo.SelectUser(search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	if ok, err := s.Model.HasPermission(uuid.MustParse(req.CallerId), usr.AdminDivID); err != nil || !ok {
		return nil, e.ErrAuthNoPermission
	}

	if err = s.Repo.UpdateUser(search, &user.User{
		IsActive: &req.Value,
	}); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &pb.PostUserActiveResponse_Data{}, nil
}

func (s *UserService) ListUsers(req *pb.GetUsersRequest) (*pb.GetUsersResponse_Data, error) {
	var err error
	var search user.Search
	var limit, skip int
	var total *int64
	limit = 10
	skip = 0

	if req.AdminDivId != "" {
		uid, err := uuid.Parse(req.AdminDivId)
		if err != nil {
			return nil, e.ErrIdInvalid
		}
		search.AdminDivID = uid
	}
	if req.Id != "" {
		uid, err := uuid.Parse(req.Id)
		if err != nil {
			return nil, e.ErrIdInvalid
		}
		search.ID = uid
	}
	if req.RoleId != "" {
		uid, err := uuid.Parse(req.RoleId)
		if err != nil {
			return nil, e.ErrIdInvalid
		}
		search.RoleID = uid
	}
	if req.Limit != "" {
		if limit, err = strconv.Atoi(req.Limit); err != nil {
			return nil, e.ErrBodyInvalid
		}
	}
	search.Limit = limit

	if req.Offset != "" {
		if skip, err = strconv.Atoi(req.Offset); err != nil {
			return nil, e.ErrBodyInvalid
		}
	}
	search.Skip = skip

	if req.Username != "" {
		search.Username = &req.Username
	}

	if total, err = s.Repo.CountUser(&search); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	users, err := s.Repo.ListUser(&search)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.GetUsersResponse_Data{
		Results:    lib.ConvertUsers(users, s.Repo),
		Pagination: lib.ConvertPagination(skip, limit, *total),
	}, nil
}

func (s *UserService) Issue(req *pb.PostUserIssueRequest) (*pb.PostUserIssueResponse_Data, error) {
	var err error
	if f, ok := validate.RequiredFields(req, "AdminDivId", "RoleId", "Id"); !ok {
		return nil, e.ErrMissingField(f)
	}
	var aid, rid, uid uuid.UUID
	if aid, err = uuid.Parse(req.AdminDivId); err != nil {
		return nil, e.ErrIdInvalid
	}
	if rid, err = uuid.Parse(req.RoleId); err != nil {
		return nil, e.ErrIdInvalid
	}
	if uid, err = uuid.Parse(req.Id); err != nil {
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

	if ok, err := s.Model.HasPermission(uid, add.ID); err == nil {
		if !ok {
			return nil, e.ErrAuthNoPermission
		}
	} else {
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

	return &pb.PostUserIssueResponse_Data{
		Username: *usr.Username,
		Password: *pass,
	}, nil
}
