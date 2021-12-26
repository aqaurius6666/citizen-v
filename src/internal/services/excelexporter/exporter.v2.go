package excelexporter

import (
	"bytes"
	"fmt"
	"io"
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
	for i, c := range ctz {
		s.handleRow(buf, i, c)
	}
	return buf, int64(buf.Len()), nil
}

func (s *ExporterV2) handleRow(writer io.Writer, i int, c *citizen.Citizen) error {
	birthday := time.UnixMilli(utils.Int64Val(c.Birthday))
	fmt.Println(birthday)
	// f.SetCellInt("Sheet1", fmt.Sprintf("A%d", i), i)
	// f.SetCellStr("Sheet1", fmt.Sprintf("B%d", i), utils.StrVal(c.Name))
	// f.SetCellStr("Sheet1", fmt.Sprintf("C%d", i), utils.StrVal(c.PID))
	// f.SetCellStr("Sheet1", fmt.Sprintf("D%d", i), utils.StrVal(c.Gender))
	// f.SetCellStr("Sheet1", fmt.Sprintf("E%d", i), utils.StrVal(c.JobName))
	// f.SetCellStr("Sheet1", fmt.Sprintf("F%d", i), utils.StrVal(c.Nationality))
	// f.SetCellStr("Sheet1", fmt.Sprintf("G%d", i), utils.StrVal(c.Religion))
	// f.SetCellStr("Sheet1", fmt.Sprintf("H%d", i), fmt.Sprintf("%d/%d/%d", birthday.Day(), birthday.Month(), birthday.Year()))
	// f.SetCellStr("Sheet1", fmt.Sprintf("J%d", i), utils.StrVal(c.EducationalLevel))
	// f.SetCellStr("Sheet1", fmt.Sprintf("K%d", i), utils.StrVal(c.Hometown))
	// f.SetCellStr("Sheet1", fmt.Sprintf("L%d", i), utils.StrVal(c.ResidencePlace))
	// f.SetCellStr("Sheet1", fmt.Sprintf("M%d", i), utils.StrVal(c.CurrentPlace))
	// f.SetCellStr("Sheet1", fmt.Sprintf("N%d", i), utils.StrVal(c.FatherName))
	// f.SetCellStr("Sheet1", fmt.Sprintf("O%d", i), utils.StrVal(c.FatherPID))
	// f.SetCellStr("Sheet1", fmt.Sprintf("P%d", i), utils.StrVal(c.MotherName))
	// f.SetCellStr("Sheet1", fmt.Sprintf("Q%d", i), utils.StrVal(c.MotherPID))
	return nil
}
