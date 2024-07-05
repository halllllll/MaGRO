package entity

type ResultLabel string

const (
	OK ResultLabel = "success"
	ER ResultLabel = "error"
)

// type Result struct {
// 	Status ResultLabel `json:"status"`
// }

type OperatorSubunits struct {
	Operators []*UserID `json:"operators"`
	Subunit   *SubUnit  `json:"subunit"`
}

type UserWithSubgroups struct {
	User
	BelongsSubunit []SubunitId
}

type RespUsersSubunit struct {
	Result        ResultLabel          `json:"status"`
	Unit          *Unit                `json:"unit"`
	CurrentUser   *User                `json:"current_user"`
	UserCount     int                  `json:"user_count"`
	UserGroups    []*UserWithSubgroups `json:"user_groups"`
	SubunitCount  int                  `json:"subunit_count"`
	SubunitGroups []*OperatorSubunits  `json:"subunit_groups"`
}

type RespBelongUnits struct {
	Result    ResultLabel `json:"status"`
	Units     []*Unit     `json:"units"`
	UnitCount int         `json:"unit_count"`
}
