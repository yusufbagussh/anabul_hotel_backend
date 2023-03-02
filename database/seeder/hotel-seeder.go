package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func HotelSeeder(province *[]entity.Province, city *[]entity.City, district *[]entity.District) *[]entity.Hotel {
	//Must userSatori UUID
	//UUIDType, errConv := uuid.FromString("69359037-9599-48e7-b8f2-48393c019135")
	//if errConv != nil {
	//	log.Fatal(errConv)
	//}

	//TO convert Raw Byte ke UUID
	//s := `"\235\273\'\021\003\261@\022\226\275o\265\322\002\211\263"`
	//s = strings.ReplaceAll(s, `\'`, `'`)
	//s2, errConv3 := strconv.Unquote(s)
	//if errConv3 != nil {
	//	fmt.Println(errConv3)
	//}
	//
	//by := []byte(s2)
	//u, errConv2 := uuid.FromBytes(by)
	//if errConv2 != nil {
	//	fmt.Println(errConv2)
	//}
	//fmt.Println(u.String())

	//To convert UUID string to UUID type
	//UUIDType, errConv := uuid.Parse("69359037-9599-48e7-b8f2-48393c019135")
	//if errConv != nil {
	//	log.Fatal(errConv)
	//}

	hotel1 := entity.Hotel{
		Name:       "Pet Hotel Indonesia",
		Email:      "pethotelindonesia@gmail.com",
		ProvinceID: (*province)[0].IDProvince,
		CityID:     (*city)[0].IDCity,
		DistrictID: (*district)[0].IDDistrict,
	}
	hotels := []entity.Hotel{
		hotel1,
	}
	return &hotels
}
