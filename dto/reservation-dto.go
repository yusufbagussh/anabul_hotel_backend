package dto

type CreateReservation struct {
	HotelID           string  `json:"hotel_id" form:"hotel_id" binding:"required"`
	UserID            string  `json:"user_id" form:"user_id" binding:"required"`
	StartDate         string  `json:"start_date" form:"start_date"`
	EndDate           string  `json:"end_date" form:"end_date"`
	TotalCost         float64 `json:"total_cost" form:"total_cost"`
	DPCost            float64 `json:"dp_cost" form:"dp_cost"`
	PaymentStatus     string  `json:"payment_status" form:"payment_status"`
	CheckInStatus     string  `json:"check_in_status" form:"check_in_status"`
	ReservationStatus string  `json:"reservation_status" form:"reservation_status"`
	//CreatedBy         string                    `json:"created_by" form:"created_by"`
	//UpdatedBy         string                    `json:"updated_by" form:"updated_by"`
	ReservationDetail []CreateReservationDetail `json:"reservation_detail"`
}

type UpdatePaymentStatus struct {
	IDReservation string `json:"id_reservation" form:"id_reservation" binding:"required"`
	PaymentStatus string `json:"payment_status" form:"payment_status" binding:"required"`
}
type UpdateCheckInStatus struct {
	IDReservation string `json:"id_reservation" form:"id_reservation" binding:"required"`
	CheckInStatus string `json:"check_in_status" form:"check_in_status" binding:"required"`
}
type UpdateReservationStatus struct {
	IDReservation     string `json:"id_reservation" form:"id_reservation" binding:"required"`
	ReservationStatus string `json:"reservation_status" form:"reservation_status" binding:"required"`
}

type UpdateReservation struct {
	IDReservation     string  `json:"id_reservation" form:"id_reservation" binding:"required"`
	HotelID           string  `json:"hotel_id" form:"hotel_id" binding:"required"`
	UserID            string  `json:"user_id" form:"user_id" binding:"required"`
	StartDate         string  `json:"start_date" form:"start_date"`
	EndDate           string  `json:"end_date" form:"end_date"`
	TotalCost         float64 `json:"total_cost" form:"total_cost"`
	DPCost            float64 `json:"dp_cost" form:"dp_cost"`
	PaymentStatus     string  `json:"payment_status" form:"payment_status" binding:"required"`
	CheckInStatus     string  `json:"check_in_status" form:"check_in_status" binding:"required"`
	ReservationStatus string  `json:"reservation_status" form:"reservation_status" binding:"required"`
	//CreatedBy         string                    `json:"created_by" form:"created_by" binding:"required"`
	//UpdatedBy         string                    `json:"updated_by" form:"updated_by" binding:"required"`
	ReservationDetail []UpdateReservationDetail `json:"reservation_detail"`
}
