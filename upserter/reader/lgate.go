package service_reader

import (
	"context"
	"encoding/csv"
	"io"

	model "github.com/halllllll/MaGRO/kajiki/model/edu"
	"github.com/shigetaichi/xsv"
)

func (r Reader) Lgate(ctx context.Context, f io.Reader) error {
	target, err := r.lgateCSVReader(f)
	if err != nil {
		return err
	}
	return r.Upsert.LgateUpsert(ctx, target)
}

func (r *Reader) lgateCSVReader(f io.Reader) ([]*model.LGateCSVOutput, error) {
	csvReader := csv.NewReader(f)
	xsvRead := xsv.NewXsvRead[*model.LGateCSVOutput]()
	xsvRead.TagName = "csv"
	xsvRead.ShouldAlignDuplicateHeadersWithStructFieldOrder = true
	xsvRead.FailIfDoubleHeaderNames = false
	xsvRead.FailIfUnmatchedStructTags = true
	var output []*model.LGateCSVOutput
	if err := xsvRead.SetReader(csvReader).ReadTo(&output); err != nil {
		return nil, err
	}
	return output, nil
}
