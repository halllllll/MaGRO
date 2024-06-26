package entity

type OperatorSubunits struct {
	Operators []*User  `json:"operators"`
	Subunit   *SubUnit `json:"subunit"`
}

type RespUsersSubunit struct {
	Unit          *Unit               `json:"unit"`
	CurrentUser   *User               `json:"current_user "`
	UserCount     int                 `json:"user_count"`
	UserGroups    []*User             `json:"user_groups"`
	SubunitCount  int                 `json:"subunit_count"`
	SubunitGroups []*OperatorSubunits `json:"subunit_groups"`
}

type RespBelongUnits struct{
	
}
