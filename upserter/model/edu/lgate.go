package edu

import (
	"fmt"
	"reflect"
	"slices"
)

type LGateRole string

const (
	RoleAdmin   LGateRole = "学校管理者"
	RoleTeacher LGateRole = "教員"
	RoleStudent LGateRole = "児童生徒"
)

var lGateRoleMap = make(map[LGateRole]string)

type LGateCSVOutput struct {
	Uuid           string `csv:"UUID"`
	Username       string `csv:"username"`
	SchoolName     string `csv:"schoolName"`
	FamilyName     string `csv:"familyName"`
	GivenName      string `csv:"givenName"`
	FamilyKanaName string `csv:"familyKanaName"`
	GivenKanaName  string `csv:"givenKanaName"`

	// 以下のやつがたくさんある(おそらく最大所属数のアカウントに引っ張られる)
	TermName1  string `csv:"termName"`
	ClassName1 string `csv:"className"`
	ClassRole1 string `csv:"classRole"`

	TermName2  string `csv:"termName"`
	ClassName2 string `csv:"className"`
	ClassRole2 string `csv:"classRole"`

	TermName3  string `csv:"termName"`
	ClassName3 string `csv:"className"`
	ClassRole3 string `csv:"classRole"`

	TermName4  string `csv:"termName"`
	ClassName4 string `csv:"className"`
	ClassRole4 string `csv:"classRole"`

	TermName5  string `csv:"termName"`
	ClassName5 string `csv:"className"`
	ClassRole5 string `csv:"classRole"`

	TermName6  string `csv:"termName"`
	ClassName6 string `csv:"className"`
	ClassRole6 string `csv:"classRole"`

	TermName7  string `csv:"termName"`
	ClassName7 string `csv:"className"`
	ClassRole7 string `csv:"classRole"`

	TermName8  string `csv:"termName"`
	ClassName8 string `csv:"className"`
	ClassRole8 string `csv:"classRole"`

	TermName9  string `csv:"termName"`
	ClassName9 string `csv:"className"`
	ClassRole9 string `csv:"classRole"`

	TermName10  string `csv:"termName"`
	ClassName10 string `csv:"className"`
	ClassRole10 string `csv:"classRole"`

	TermName11  string `csv:"termName"`
	ClassName11 string `csv:"className"`
	ClassRole11 string `csv:"classRole"`

	TermName12  string `csv:"termName"`
	ClassName12 string `csv:"className"`
	ClassRole12 string `csv:"classRole"`

	TermName13  string `csv:"termName"`
	ClassName13 string `csv:"className"`
	ClassRole13 string `csv:"classRole"`

	TermName14  string `csv:"termName"`
	ClassName14 string `csv:"className"`
	ClassRole14 string `csv:"classRole"`

	TermName15  string `csv:"termName"`
	ClassName15 string `csv:"className"`
	ClassRole15 string `csv:"classRole"`

	TermName16  string `csv:"termName"`
	ClassName16 string `csv:"className"`
	ClassRole16 string `csv:"classRole"`

	TermName17  string `csv:"termName"`
	ClassName17 string `csv:"className"`
	ClassRole17 string `csv:"classRole"`

	TermName18  string `csv:"termName"`
	ClassName18 string `csv:"className"`
	ClassRole18 string `csv:"classRole"`

	TermName19  string `csv:"termName"`
	ClassName19 string `csv:"className"`
	ClassRole19 string `csv:"classRole"`

	TermName20  string `csv:"termName"`
	ClassName20 string `csv:"className"`
	ClassRole20 string `csv:"classRole"`

	TermName21  string `csv:"termName"`
	ClassName21 string `csv:"className"`
	ClassRole21 string `csv:"classRole"`

	TermName22  string `csv:"termName"`
	ClassName22 string `csv:"className"`
	ClassRole22 string `csv:"classRole"`

	TermName23  string `csv:"termName"`
	ClassName23 string `csv:"className"`
	ClassRole23 string `csv:"classRole"`

	TermName24  string `csv:"termName"`
	ClassName24 string `csv:"className"`
	ClassRole24 string `csv:"classRole"`

	TermName25  string `csv:"termName"`
	ClassName25 string `csv:"className"`
	ClassRole25 string `csv:"classRole"`

	TermName26  string `csv:"termName"`
	ClassName26 string `csv:"className"`
	ClassRole26 string `csv:"classRole"`

	TermName27  string `csv:"termName"`
	ClassName27 string `csv:"className"`
	ClassRole27 string `csv:"classRole"`

	TermName28  string `csv:"termName"`
	ClassName28 string `csv:"className"`
	ClassRole28 string `csv:"classRole"`

	TermName29  string `csv:"termName"`
	ClassName29 string `csv:"className"`
	ClassRole29 string `csv:"classRole"`

	TermName30  string `csv:"termName"`
	ClassName30 string `csv:"className"`
	ClassRole30 string `csv:"classRole"`

	TermName31  string `csv:"termName"`
	ClassName31 string `csv:"className"`
	ClassRole31 string `csv:"classRole"`

	TermName32  string `csv:"termName"`
	ClassName32 string `csv:"className"`
	ClassRole32 string `csv:"classRole"`

	TermName33  string `csv:"termName"`
	ClassName33 string `csv:"className"`
	ClassRole33 string `csv:"classRole"`

	TermName34  string `csv:"termName"`
	ClassName34 string `csv:"className"`
	ClassRole34 string `csv:"classRole"`

	TermName35  string `csv:"termName"`
	ClassName35 string `csv:"className"`
	ClassRole35 string `csv:"classRole"`

	TermName36  string `csv:"termName"`
	ClassName36 string `csv:"className"`
	ClassRole36 string `csv:"classRole"`

	TermName37  string `csv:"termName"`
	ClassName37 string `csv:"className"`
	ClassRole37 string `csv:"classRole"`

	TermName38  string `csv:"termName"`
	ClassName38 string `csv:"className"`
	ClassRole38 string `csv:"classRole"`

	TermName39  string `csv:"termName"`
	ClassName39 string `csv:"className"`
	ClassRole39 string `csv:"classRole"`

	TermName40  string `csv:"termName"`
	ClassName40 string `csv:"className"`
	ClassRole40 string `csv:"classRole"`

	TermName41  string `csv:"termName"`
	ClassName41 string `csv:"className"`
	ClassRole41 string `csv:"classRole"`

	TermName42  string `csv:"termName"`
	ClassName42 string `csv:"className"`
	ClassRole42 string `csv:"classRole"`

	TermName43  string `csv:"termName"`
	ClassName43 string `csv:"className"`
	ClassRole43 string `csv:"classRole"`

	TermName44  string `csv:"termName"`
	ClassName44 string `csv:"className"`
	ClassRole44 string `csv:"classRole"`

	TermName45  string `csv:"termName"`
	ClassName45 string `csv:"className"`
	ClassRole45 string `csv:"classRole"`

	TermName46  string `csv:"termName"`
	ClassName46 string `csv:"className"`
	ClassRole46 string `csv:"classRole"`

	TermName47  string `csv:"termName"`
	ClassName47 string `csv:"className"`
	ClassRole47 string `csv:"classRole"`

	TermName48  string `csv:"termName"`
	ClassName48 string `csv:"className"`
	ClassRole48 string `csv:"classRole"`

	TermName49  string `csv:"termName"`
	ClassName49 string `csv:"className"`
	ClassRole49 string `csv:"classRole"`

	TermName50  string `csv:"termName"`
	ClassName50 string `csv:"className"`
	ClassRole50 string `csv:"classRole"`
}

func (l *LGateCSVOutput) RowSubunits() []string {
	lGateRoleMap[RoleAdmin] = "teacher"
	lGateRoleMap[RoleTeacher] = "teacher"
	lGateRoleMap[RoleStudent] = "student"

	var result []string
	val := reflect.ValueOf(l).Elem()
	for i := 1; i <= 50; i++ {
		termField := val.FieldByName(fmt.Sprintf("TermName%d", i))
		classField := val.FieldByName(fmt.Sprintf("ClassName%d", i))
		if termField.IsValid() && classField.IsValid() {
			term := termField.String()
			class := classField.String()
			if term != "" && class != "" {
				result = append(result, fmt.Sprintf("%s:%s", term, class))
			}
		}
	}
	return result
}

func (l *LGateCSVOutput) RowRoles() []LGateRole {
	var result []LGateRole
	val := reflect.ValueOf(l).Elem()
	for i := 1; i <= 50; i++ {
		roleField := val.FieldByName(fmt.Sprintf("ClassRole%d", i))
		if roleField.IsValid() {
			role := roleField.String()
			if role != "" {
				result = append(result, LGateRole(role))
			}
		}
	}
	return result
}

func (l *LGateCSVOutput) IsStudent() bool {
	return slices.Contains(l.RowRoles(), RoleStudent)
}
