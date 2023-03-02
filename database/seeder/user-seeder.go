package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func UserSeeder(hotel *[]entity.Hotel) *[]entity.User {
	hash, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	user1 := entity.User{
		Name:     "Hashirama Senju",
		Email:    "hashirama@gmail.com",
		Password: string(hash),
		Role:     "Super Admin",
		Verified: true,
	}
	user2 := entity.User{
		Name:     "Tobirama Senju",
		Email:    "tobirama@gmail.com",
		Password: string(hash),
		Role:     "Admin",
		HotelID:  (*hotel)[0].IDHotel,
		Verified: true,
	}
	user3 := entity.User{
		Name:     "Hiruzen Sarutobi",
		Email:    "hiruzen@gmail.com",
		Password: string(hash),
		Role:     "Staff",
		HotelID:  (*hotel)[0].IDHotel,
		Verified: true,
	}
	user4 := entity.User{
		Name:     "Namikaze Minato",
		Email:    "minato@gmail.com",
		NIK:      44112238478324328,
		Password: string(hash),
		Role:     "Customer",
		Verified: true,
	}
	users := &[]entity.User{
		user1,
		user2,
		user3,
		user4,
	}
	return users
}
