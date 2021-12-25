package excelexporter

import (
	"io"

	"github.com/aqaurius6666/citizen-v/src/internal/db/citizen"
)

type Exporter interface {
	ExportCitizen([]*citizen.Citizen, io.Writer) error
}
