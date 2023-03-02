package dto

// ResetPasswordInput struct
type ResetPasswordInput struct {
	Email           string `json:"email" form:"email" binding:"required,email" `
	Token           string `json:"token" form:"token" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}
