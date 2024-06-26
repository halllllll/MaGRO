package model

import "time"

type MaGRORoleID int
type MaGRORoleName string
type MaGROStatus string

type MaGROUserID string
type MaGROAccountID string

type MaGROUnitID int
type MaGROSubunitID int

type MaGROSubUnitID int

type MaGROStatusID int

type MaGROActionID int

const (
	StatusActive  MaGROStatus   = "active"
	StatusSuspend MaGROStatus   = "suspend"
	RoleAdmin     MaGRORoleName = "admin"
	RoleDirector  MaGRORoleName = "director"

	RoleManager MaGRORoleName = "manager"
	RoleGuest   MaGRORoleName = "guest"
)

// each tables
type MaGROUsersStatus struct {
	Id   MaGROStatusID
	Name string
}

type MaGRORole struct {
	Id   MaGRORoleID
	Name MaGRORoleName
}

type MaGROMiddleTableUsersUnit struct {
	Id        MaGROUserID
	AccountId MaGROAccountID
}

type MaGROUnit struct {
	Id   MaGROUnitID
	Name string
}

type MaGROSubUnit struct {
	Id     MaGROSubUnitID
	UnitId MaGROUnitID
	Name   string
	Public bool
}

type MaGROMiddleTableUsersSubunit struct {
	Id        int
	UserId    MaGROUserID
	SubUnitId MaGROSubUnitID
}

type MaGROUser struct {
	Id        MaGROUserID
	AccountId MaGROAccountID
	Name      string
	Kana      string
	RoleId    MaGRORoleID
	Status    MaGROStatusID
}

type MaGROAction struct {
	Id   MaGROActionID
	Name string
}

type MaGROLogs struct {
	Id        int
	Timestamp time.Time
	UserId    MaGROUserID
	ActionId  MaGROActionID
}
