package seeder

//
//import (
//	"github.com/yusufbagussh/pet_hotel_backend/entity"
//	"github.com/yusufbagussh/pet_hotel_backend/utils"
//	"gorm.io/gorm"
//)
//
//type Seeder interface {
//	Seeder()
//}
//
//type seederConnection struct {
//	db *gorm.DB
//}
//
//func (s seederConnection) Seeder() {
//	if utils.EnvVar("APP_ENV", "DEVELOPMENT") == "DEVELOPMENT" {
//		province := entity.Province{
//			Name: "Jawa Tengah",
//		}
//		s.db.Create(&province)
//		city := entity.City{
//			ProvinceID: province.IDProvince,
//			Name:       "Surakarta",
//		}
//		s.db.Create(&city)
//		district := entity.District{
//			CityID: city.IDCity,
//			Name:   "Jebres",
//		}
//		district2 := entity.District{
//			CityID: city.IDCity,
//			Name:   "Nusukan",
//		}
//		district3 := entity.District{
//			CityID: city.IDCity,
//			Name:   "Mojosongo",
//		}
//		districts := &[]entity.District{
//			district,
//			district2,
//			district3,
//		}
//		s.db.Create(&districts)
//	}
//}
//
//func NewSeeder(db *gorm.DB) Seeder {
//	return &seederConnection{db: db}
//}
