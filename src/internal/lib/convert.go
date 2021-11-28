package lib

import (
	"github.com/aquarius6666/citizen-v/src/internal/db/admindiv"
	"github.com/aquarius6666/citizen-v/src/pb"
)

func ConvertAdminDiv(db []*admindiv.AdminDiv) []*pb.AdminDiv {
	r := make([]*pb.AdminDiv, 0)
	for _, s := range db {
		r = append(r, &pb.AdminDiv{
			Code:       *s.Code,
			Name:       *s.Name,
			SuperiorId: s.SuperiorID.String(),
			Type:       *s.Type,
			Id:         s.ID.String(),
			Subdiv:     ConvertAdminDiv(s.SubDiv),
		})
	}
	return r
}
