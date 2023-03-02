package helper

import (
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

type CheckHelper interface {
	CheckHotel(hotelID string, userID string) bool
	CheckAdminHotel(userID string) bool
}

type checkHelper struct {
	userRepository repository.UserRepository
}

func NewCheckHelper(userRepo repository.UserRepository) CheckHelper {
	return &checkHelper{
		userRepository: userRepo,
	}
}

func (c *checkHelper) CheckHotel(hotelID string, userID string) bool {
	result, _ := c.userRepository.FindUserByID(userID)
	if hotelID == result.HotelID {
		return true
	} else {
		return false
	}
}

func (c *checkHelper) CheckAdminHotel(userID string) bool {
	result, _ := c.userRepository.FindUserByID(userID)
	if result.Role == "Admin" {
		return true
	}
	return false
}
