package excelexporter

import (
	"io"

	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
	"github.com/xuri/excelize/v2"
	"golang.org/x/xerrors"
)

var (
	_ Exporter = (*ExporterV1)(nil)
)

type ExporterV1 struct {
}

func (s *ExporterV1) ExportCitizen(ctz []*citizen.Citizen, writer io.Writer) error {
	f := excelize.NewFile()
	f.NewSheet("test")
	out, err := f.NewStreamWriter("test")
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	out.File.WriteTo(writer)
	return nil
}
