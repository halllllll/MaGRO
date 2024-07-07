package entity

type UserID string
type UserUUID string

type RespCurrentUser struct {
	UserID   UserID `json:"user_id"`
	UserName string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}

type User struct {
	UserID      UserUUID `json:"user_id"`
	UserName    UserID   `json:"user_name"`
	DisplayName string   `json:"user_displayname"`
	UserSortKey string   `json:"user_sortkey"`
	UserType    Role     `json:"user_role"`
	Status      string   `json:"user_status"`
}
