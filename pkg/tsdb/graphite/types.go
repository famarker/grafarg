package graphite

import "github.com/famarker/grafarg/pkg/tsdb"

type TargetResponseDTO struct {
	Target     string                `json:"target"`
	DataPoints tsdb.TimeSeriesPoints `json:"datapoints"`
}
