package dto

type Device struct {
	UserID      string `json:"user_id" form:"user_id"`
	DeviceToken string `json:"device_token" form:"user_id"`
}
