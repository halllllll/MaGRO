package entity

import "time"

type SystemVersion string

type System struct {
	Version  SystemVersion `json:"version" db:"version"`
	Created  time.Time     `json:"created" db:"created_at"`
	Modified time.Time     `json:"modified" db:"updated_at"`
}

/* API request body */

type ReqNewRoleAlias struct {
	AliasAdminName    string `json:"admin_alias"`
	AliasDirectorName string `json:"director_alias"`
	AliasManagerName  string `json:"manager_alias"`
	AliasGuestName    string `json:"guest_alias"`
}
