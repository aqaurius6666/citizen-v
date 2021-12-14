package api

import (
	"strconv"

	"github.com/aqaurius6666/citizen-v/src/internal/db"
	"github.com/aqaurius6666/citizen-v/src/internal/db/user"
	"github.com/aqaurius6666/citizen-v/src/internal/lib"
	"github.com/aqaurius6666/citizen-v/src/internal/var/e"
	"github.com/aqaurius6666/citizen-v/src/pb"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type UserService struct {
	Repo db.ServerRepo
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
