package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"strings"
)

// HotelRepository is contract what hotelRepository can do to db
type HotelRepository interface {
	GetAllHotel(filterPagination dto.FilterPagination) ([]entity.Hotel, dto.Pagination, error)
	InsertHotel(hotel entity.Hotel) (entity.Hotel, error)
	InsertHotelAdmin(hotel entity.Hotel, user entity.User) (entity.Hotel, error)
	UpdateHotel(hotel entity.Hotel) (entity.Hotel, error)
	FindHotelByID(hotelID string) (entity.Hotel, error)
	DeleteHotel(hotel entity.Hotel) error
}

// hotelConnection adalah func untuk melakukan query data ke tabel hotel
type hotelConnection struct {
	connection *gorm.DB
}

func (db *hotelConnection) DeleteHotel(hotel entity.Hotel) error {
	err := db.connection.Where("id_hotel = ?", hotel.IDHotel).Delete(&hotel).Error
	return err
}

func (db *hotelConnection) GetAllHotel(filterPagination dto.FilterPagination) ([]entity.Hotel, dto.Pagination, error) {
	search := filterPagination.Search
	sortBy := filterPagination.SortBy
	orderBy := filterPagination.OrderBy
	perPage := int(filterPagination.PerPage)
	page := int(filterPagination.Page)

	if sortBy == "" {
		sortBy = "created_at"
	}
	if orderBy == "" {
		orderBy = "desc"
	}
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	var total int64

	var hotels []entity.Hotel

	query := db.connection.Model(&hotels).Joins("LEFT JOIN provinces ON hotels.province_id = provinces.id_province").
		Joins("LEFT JOIN cities ON hotels.city_id = cities.id_city").
		Joins("LEFT JOIN districts ON hotels.district_id = districts.id_district")

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				query = db.Where("LOWER(hotels.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(provinces.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(cities.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(districts.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return query
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.ProvinceID != "" {
			db.Where("hotels.province_id = ?", filterPagination.ProvinceID)
		}
		if filterPagination.CityID != "" {
			db.Where("hotels.city_id = ?", filterPagination.CityID)
		}
		if filterPagination.DistrictID != "" {
			db.Where("hotels.district_id = ?", filterPagination.DistrictID)
		}
		return db
	})

	if sortBy != "" && orderBy != "" {
		query.Order(
			clause.OrderByColumn{
				Column: clause.Column{
					Name: strings.ToLower(sortBy),
				},
				Desc: orderBy == "DESC",
			},
		)
	}

	//query.Find(&hotels).Count(&total)

	err := query.Count(&total).Limit(perPage).Offset((page - 1) * perPage).
		Preload("Province").
		Preload("City").
		Preload("District").
		Find(&hotels).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}
	return hotels, pagination, err
}

// InsertHotelAdmin is to add hotel in database
func (db *hotelConnection) InsertHotelAdmin(hotel entity.Hotel, user entity.User) (entity.Hotel, error) {
	var err error
	err = db.connection.Save(&hotel).Error
	user.HotelID = hotel.IDHotel
	err = db.connection.Save(&user).Error
	return hotel, err
}

func (db *hotelConnection) InsertHotel(hotel entity.Hotel) (entity.Hotel, error) {
	var err error
	err = db.connection.Save(&hotel).Error
	return hotel, err
}

// UpdateHotel is func to edit hotel in database
func (db *hotelConnection) UpdateHotel(hotel entity.Hotel) (entity.Hotel, error) {
	err := db.connection.Where("id_hotel = ?", hotel.IDHotel).Updates(&hotel).Error
	return hotel, err
}

// FindHotelByID is func to get hotel by email
func (db *hotelConnection) FindHotelByID(hotelID string) (entity.Hotel, error) {
	var hotel entity.Hotel
	err := db.connection.Where("id_hotel = ?", hotelID).Preload("Province").Preload("City").Preload("District").Take(&hotel).Error
	return hotel, err
}

// NewHotelRepository is creates a new instance of HotelRepository
func NewHotelRepository(db *gorm.DB) HotelRepository {
	return &hotelConnection{
		connection: db,
	}
}
