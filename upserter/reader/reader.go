package service_reader

import (
	"github.com/halllllll/MaGRO/kajiki/upsert"
)

type Reader struct {
	Upsert upsert.Upsert
}

type Format string

const (
	FormatLGate      Format = "lgate"
	_FormatLoilo     Format = "loilo"
	_FormatC4th      Format = "c4th"
	_FormatMiraiseed Format = "miraiseed"
)

func NewReader(u *upsert.Upsert) *Reader {
	return &Reader{Upsert: *u}
}
