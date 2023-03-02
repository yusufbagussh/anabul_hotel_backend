package migration

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
)

type Migrator interface {
	DropTable()
	Migration()
}

type migrationConnection struct {
	db *gorm.DB
}

func (m migrationConnection) DropTable() {
	if m.db.Migrator().HasTable(&entity.Province{}) {
		m.db.Migrator().DropTable(&entity.Province{})
	}
	if m.db.Migrator().HasTable(&entity.City{}) {
		m.db.Migrator().DropTable(&entity.City{})
	}
	if m.db.Migrator().HasTable(&entity.District{}) {
		m.db.Migrator().DropTable(&entity.District{})
	}
	if m.db.Migrator().HasTable(&entity.Request{}) {
		m.db.Migrator().DropTable(&entity.Request{})
	}
	if m.db.Migrator().HasTable(&entity.Hotel{}) {
		m.db.Migrator().DropTable(&entity.Hotel{})
	}
	if m.db.Migrator().HasTable(&entity.HotelAlbum{}) {
		m.db.Migrator().DropTable(&entity.HotelAlbum{})
	}
	if m.db.Migrator().HasTable(&entity.User{}) {
		m.db.Migrator().DropTable(&entity.User{})
	}
	if m.db.Migrator().HasTable(&entity.Class{}) {
		m.db.Migrator().DropTable(&entity.Class{})
	}
	if m.db.Migrator().HasTable(&entity.Category{}) {
		m.db.Migrator().DropTable(&entity.Category{})
	}
	if m.db.Migrator().HasTable(&entity.CageCategory{}) {
		m.db.Migrator().DropTable(&entity.CageCategory{})
	}
	if m.db.Migrator().HasTable(&entity.CageType{}) {
		m.db.Migrator().DropTable(&entity.CageType{})
	}
	if m.db.Migrator().HasTable(&entity.CageDetail{}) {
		m.db.Migrator().DropTable(&entity.CageDetail{})
	}
	if m.db.Migrator().HasTable(&entity.Cage{}) {
		m.db.Migrator().DropTable(&entity.Cage{})
	}
	if m.db.Migrator().HasTable(&entity.Pet{}) {
		m.db.Migrator().DropTable(&entity.Pet{})
	}
	if m.db.Migrator().HasTable(&entity.Product{}) {
		m.db.Migrator().DropTable(&entity.Product{})
	}
	if m.db.Migrator().HasTable(&entity.Species{}) {
		m.db.Migrator().DropTable(&entity.Species{})
	}
	if m.db.Migrator().HasTable(&entity.Group{}) {
		m.db.Migrator().DropTable(&entity.Group{})
	}
	if m.db.Migrator().HasTable(&entity.GroupDetail{}) {
		m.db.Migrator().DropTable(&entity.GroupDetail{})
	}
	if m.db.Migrator().HasTable(&entity.Group{}) {
		m.db.Migrator().DropTable(&entity.Group{})
	}
	if m.db.Migrator().HasTable(&entity.Service{}) {
		m.db.Migrator().DropTable(&entity.Service{})
	}
	if m.db.Migrator().HasTable(&entity.ServiceDetail{}) {
		m.db.Migrator().DropTable(&entity.ServiceDetail{})
	}
	if m.db.Migrator().HasTable(&entity.Reservation{}) {
		m.db.Migrator().DropTable(&entity.Reservation{})
	}
	if m.db.Migrator().HasTable(&entity.ReservationDetail{}) {
		m.db.Migrator().DropTable(&entity.ReservationDetail{})
	}
	if m.db.Migrator().HasTable(&entity.ReservationProduct{}) {
		m.db.Migrator().DropTable(&entity.ReservationProduct{})
	}
	if m.db.Migrator().HasTable(&entity.ReservationService{}) {
		m.db.Migrator().DropTable(&entity.ReservationService{})
	}
	if m.db.Migrator().HasTable(&entity.ReservationInventory{}) {
		m.db.Migrator().DropTable(&entity.ReservationInventory{})
	}
	if m.db.Migrator().HasTable(&entity.ReservationCondition{}) {
		m.db.Migrator().DropTable(&entity.ReservationCondition{})
	}
	if m.db.Migrator().HasTable(&entity.ReservationConditionAlbum{}) {
		m.db.Migrator().DropTable(&entity.ReservationConditionAlbum{})
	}
	if m.db.Migrator().HasTable(&entity.Rate{}) {
		m.db.Migrator().DropTable(&entity.Rate{})
	}
	if m.db.Migrator().HasTable(&entity.Response{}) {
		m.db.Migrator().DropTable(&entity.Response{})
	}
}

func (m migrationConnection) Migration() {
	m.db.AutoMigrate(
		entity.Province{},
		entity.City{},
		entity.District{},
		entity.Request{},
		entity.Hotel{},
		entity.HotelAlbum{},
		entity.User{},
		entity.Class{},
		entity.Category{},
		entity.Species{},
		entity.Pet{},
		entity.CageCategory{},
		entity.CageType{},
		entity.CageDetail{},
		entity.Cage{},
		entity.Product{},
		entity.Group{},
		entity.GroupDetail{},
		entity.Service{},
		entity.ServiceDetail{},
		entity.Reservation{},
		entity.ReservationDetail{},
		entity.ReservationProduct{},
		entity.ReservationService{},
		entity.ReservationInventory{},
		entity.ReservationCondition{},
		entity.ReservationConditionAlbum{},
		entity.Rate{},
		entity.Response{},
	)
}

func NewMigration(conn *gorm.DB) Migrator {
	return &migrationConnection{
		db: conn,
	}
}
