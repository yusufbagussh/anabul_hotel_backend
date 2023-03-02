package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ServiceDetailSeeder(groups *[]entity.Group, services *[]entity.Service, hotels *[]entity.Hotel) *[]entity.ServiceDetail {
	groupDetail1 := entity.ServiceDetail{
		GroupID:   (*groups)[0].IDGroup,
		ServiceID: (*services)[0].IDService,
		HotelID:   (*hotels)[0].IDHotel,
		Price:     75000,
	}
	groupDetail2 := entity.ServiceDetail{
		GroupID:   (*groups)[0].IDGroup,
		ServiceID: (*services)[1].IDService,
		HotelID:   (*hotels)[0].IDHotel,
		Price:     75000,
	}
	groupDetails := []entity.ServiceDetail{
		groupDetail1,
		groupDetail2,
	}
	return &groupDetails
}
