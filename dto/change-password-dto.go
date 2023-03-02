package dto

// ChangePasswordDTO is used when client post from /changepassword url
type ChangePasswordDTO struct {
	OldPassword  string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword  string `json:"new_password" form:"new_password" binding:"required"`
	ConfPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
}
