package excelexporter

import (
	"io"

	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
)

var (
	_ Exporter = (*ExporterV1)(nil)
)

type ExporterV1 struct {
}

func (s *ExporterV1) ExportCitizen(ctz []*citizen.Citizen, writer io.Writer) error {
	return nil
}
