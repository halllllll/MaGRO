package entity

type UserID string
type UserUUID string

type RespCurrentUser struct {
	UserID   UserID `json:"user_id"`
	UserName string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}

type User struct {
	UserID      UserUUID
	UserName    UserID
	DisplayName string
	UserSortKey string
	UserType    Role
	Status      string
}
