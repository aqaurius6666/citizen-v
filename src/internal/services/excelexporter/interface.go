package excelexporter

import (
	"io"

	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
)

type Exporter interface {
	ExportCitizen([]*citizen.Citizen) (io.Reader, int64, error)
}

func NewExporter() Exporter {
	return &ExporterV2{}
}
