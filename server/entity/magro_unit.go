package entity

import "time"

type UnitId int
type SubunitId int

type Unit struct {
	UnitID UnitId `json:"unit_id"`
	Name   string `json:"name"`
}

type SubUnit struct {
	SubunitID SubunitId `json:"subunit_id"`
	Name      string    `json:"name"`
	IsPublic  bool      `json:"isPublic"`
	CreatdAt  time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}
