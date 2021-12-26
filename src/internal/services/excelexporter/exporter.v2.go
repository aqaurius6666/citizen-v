package excelexporter

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	"github.com/aqaurius6666/go-utils/utils"
)

var (
	_ Exporter = (*ExporterV2)(nil)
)

type ExporterV2 struct {
}

func (s *ExporterV2) ExportCitizen(ctz []*citizen.Citizen) (io.Reader, int64, error) {
	buf := new(bytes.Buffer)
	s.Header(buf)
	for i, c := range ctz {
		s.handleRow(buf, i, c)
	}
	return buf, int64(buf.Len()), nil
}

func (s *ExporterV2) Header(writer io.Writer) error {
	row := make([]string, 0)
	row = append(row, "STT")
	row = append(row, "Họ Tên")
	row = append(row, "CMND/CCCD")
	row = append(row, "Giới Tính")
	row = append(row, "Nghề Nghiệp")
	row = append(row, "Quốc tịch")
	row = append(row, "Tôn Giáo")
	row = append(row, "Ngày Sinh")
	row = append(row, "Trình Độ Văn Hóa")
	row = append(row, "Quê Quán")
	row = append(row, "Nơi Ở Thường Trú")
	row = append(row, "Họ Tên Bố")
	row = append(row, "CMND/CCCD Bố")
	row = append(row, "Họ Tên Mẹ")
	row = append(row, "CMND/CCCD Mẹ")

	ret := strings.Join(row, ";") + "\n"
	writer.Write([]byte(ret))
	return nil
}

func (s *ExporterV2) handleRow(writer io.Writer, i int, c *citizen.Citizen) error {
	birthday := time.UnixMilli(utils.Int64Val(c.Birthday))
	row := make([]string, 0)
	row = append(row, fmt.Sprint(i))
	row = append(row, utils.StrVal(c.Name))
	row = append(row, utils.StrVal(c.PID))
	row = append(row, utils.StrVal(c.Gender))
	row = append(row, utils.StrVal(c.JobName))
	row = append(row, utils.StrVal(c.Nationality))
	row = append(row, utils.StrVal(c.Religion))
	row = append(row, fmt.Sprintf("%d/%d/%d", birthday.Day(), birthday.Month(), birthday.Year()))
	row = append(row, utils.StrVal(c.EducationalLevel))
	row = append(row, utils.StrVal(c.Hometown))
	row = append(row, utils.StrVal(c.ResidencePlace))
	row = append(row, utils.StrVal(c.FatherName))
	row = append(row, utils.StrVal(c.FatherPID))
	row = append(row, utils.StrVal(c.MotherName))
	row = append(row, utils.StrVal(c.MotherPID))

	ret := strings.Join(row, ";") + "\n"

	writer.Write([]byte(ret))
	return nil
}
