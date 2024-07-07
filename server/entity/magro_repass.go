package entity

type users struct{
	ID UserUUID `json:"user_id" binding:"required"`
	Account UserID `json:"user_account" binding:"required"`
}
type RepassRequest struct{
	CurrentUser User `json:"current_user" binding:"required"`
	TargetUsers []users	`json:"target_users" binding:"required"`
}