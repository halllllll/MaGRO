package entity

type Role string

// Roleを表す値はDBで設定されているが
const (
	RoleAdmin    Role = "admin"
	RoleDirector Role = "director"
	RoleManager  Role = "manager"
	RoleGuest    Role = "guest"
)
