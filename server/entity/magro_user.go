package entity

type UserID string

type RespCurrentUser struct {
	UserID   UserID `json:"user_id"`
	UserName string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}

type User struct {
	UserID      UserID
	UserName    string
	DisplayName string
	UserSortKey string
	UserType    Role
}
