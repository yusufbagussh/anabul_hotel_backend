package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func RequestSeeder() *[]entity.Request {
	request1 := entity.Request{
		HotelName:  "Pet Hotel Indonesia",
		HotelEmail: "pethotelindonesia@gmail.com",
		HotelPhone: 89670198915,
		AdminName:  "Pet Hotel Indonesia",
		AdminPhone: 89670198915,
		NPWP:       "NPWP_Pet Hotel Indonesia.jpg",
		Document:   "Document_Pet Hotel Indonesia.jpg",
		NIK:        223344556677,
		KTP:        "KTP_Pet Hotel Indonesia.jpg",
		Selfie:     "Selfie_Pet Hotel Indonesia.jpg",
		Status:     "Diterima",
	}
	requests := []entity.Request{
		request1,
	}
	return &requests
}
