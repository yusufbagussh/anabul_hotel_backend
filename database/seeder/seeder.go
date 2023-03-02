package seeder

import (
	"gorm.io/gorm"
)

type Seeder interface {
	Seeder()
}

type seederConnection struct {
	db *gorm.DB
}

func (s *seederConnection) Seeder() {
	provinces := ProvinceSeeder()
	s.db.Create(provinces)
	cities := CitySeeder(provinces)
	s.db.Create(cities)
	districts := DistrictSeeder(cities)
	s.db.Create(districts)
	hotels := HotelSeeder(provinces, cities, districts)
	s.db.Create(hotels)
	users := UserSeeder(hotels)
	s.db.Create(users)
	classes := ClassSeeder()
	s.db.Create(classes)
	categories := CategorySeeder(classes)
	s.db.Create(categories)
	species := SpeciesSeeder(categories)
	s.db.Create(species)
	cageCategories := CageCategorySeeder(hotels)
	s.db.Create(cageCategories)
	cageTypes := CageTypeSeeder(hotels)
	s.db.Create(cageTypes)
	cageDetails := CageDetailSeeder(cageCategories, cageTypes, hotels)
	s.db.Create(cageDetails)
	cages := CageSeeder(cageDetails, hotels)
	s.db.Create(cages)
	pets := PetSeeder(species, users)
	s.db.Create(pets)
	products := ProductSeeder(hotels)
	s.db.Create(products)
	groups := GroupSeeder(hotels)
	s.db.Create(groups)
	groupDetails := GroupDetailSeeder(groups, species, hotels)
	s.db.Create(groupDetails)
	services := ServiceSeeder(hotels)
	s.db.Create(services)
	serviceDetails := ServiceDetailSeeder(groups, services, hotels)
	s.db.Create(serviceDetails)
	reservations := ReservationSeeder(hotels, users)
	s.db.Create(reservations)
	reservationDetails := ReservationDetailSeeder(reservations, pets, cages, products)
	s.db.Create(reservationDetails)
	reservationConditions := ReservationConditionSeeder(reservationDetails, users)
	s.db.Create(reservationConditions)
	reservationServiceDetails := ReservationServiceSeeder(reservationDetails, services)
	s.db.Create(reservationServiceDetails)
	rates := RateSeeder(reservations)
	s.db.Create(rates)
	responses := ResponseSeeder(rates)
	s.db.Create(responses)
}

func NewSeeder(conn *gorm.DB) Seeder {
	return &seederConnection{db: conn}
}

//func DBSeed(db *gorm.DB) error {
//	for _, seeder := range RegisterSeeders(db) {
//		err := db.Debug().Create(seeder.Seed).Error
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
