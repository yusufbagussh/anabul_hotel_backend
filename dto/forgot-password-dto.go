package dto

// ForgotPasswordDTO is used when client post from /changepassword url
type ForgotPasswordDTO struct {
	Email string `json:"email" form:"email" binding:"required,email" `
}
