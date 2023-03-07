package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// ReservationRepository is contract what reservationRepositoryservice.n do to db
type ReservationRepository interface {
	GetAllGroupDetail(hotelID string) ([]entity.GroupDetail, error)
	GetAllReservation(hotelID string, filterPagination dto.FilterPagination) ([]entity.Reservation, dto.Pagination, error)
	InsertReservation(reservation entity.Reservation, reservationDetail []entity.ReservationDetail, reservationServices []entity.ReservationService, reservationInventories []entity.ReservationInventory, transDetail []dto.TransactionDetail) (entity.Reservation, error)
	UpdateReservation(reservation entity.Reservation, reservationDetail []entity.ReservationDetail, reservationServices []entity.ReservationService, reservationInventories []entity.ReservationInventory) (entity.Reservation, error)
	DeleteReservation(reservation entity.Reservation) error
	FindReservationByID(reservationID string) (entity.Reservation, error)
	FindProductByID(productID string) (entity.Product, error)
	FindServiceDetailByID(serviceID string, groupID string) (entity.ServiceDetail, error)
	FindCageDetailByID(cageDetailID string) (entity.CageDetail, error)
	UpdatePaymentStatus(paymentStatus dto.UpdatePaymentStatus) (entity.Reservation, error)
	UpdateReservationStatus(reservationStatus dto.UpdateReservationStatus) (entity.Reservation, error)
	UpdateCheckInStatus(checkinStatus dto.UpdateCheckInStatus) (entity.Reservation, error)
	CreateNotification(notification entity.ReservationNotification) (entity.ReservationNotification, error)
}

// reservationConnection adalah func untuk melakukan query data ke tabel reservation
type reservationConnection struct {
	connection *gorm.DB
}

func (db *reservationConnection) CreateNotification(notification entity.ReservationNotification) (entity.ReservationNotification, error) {
	err := db.connection.Save(&notification).Error
	return notification, err
}

func (db *reservationConnection) UpdatePaymentStatus(paymentStatus dto.UpdatePaymentStatus) (entity.Reservation, error) {
	var reservation entity.Reservation
	err := db.connection.Model(&reservation).Where("id_reservation = ?", paymentStatus.IDReservation).Updates(&entity.Reservation{PaymentStatus: paymentStatus.PaymentStatus}).Error
	db.connection.Find(&reservation)
	return reservation, err
}
func (db *reservationConnection) UpdateReservationStatus(paymentStatus dto.UpdateReservationStatus) (entity.Reservation, error) {
	var reservation entity.Reservation
	err := db.connection.Model(&reservation).Where("reservation_id = ?", paymentStatus.IDReservation).Updates(&entity.Reservation{ReservationStatus: reservation.ReservationStatus}).Error
	db.connection.Find(&reservation)
	return reservation, err
}
func (db *reservationConnection) UpdateCheckInStatus(paymentStatus dto.UpdateCheckInStatus) (entity.Reservation, error) {
	var reservation entity.Reservation
	err := db.connection.Model(&reservation).Where("reservation_id = ?", paymentStatus.IDReservation).Updates(&entity.Reservation{CheckInStatus: reservation.CheckInStatus}).Error
	db.connection.Find(&reservation)
	return reservation, err
}

func (db *reservationConnection) FindServiceDetailByID(serviceID string, groupID string) (entity.ServiceDetail, error) {
	var serviceDetail entity.ServiceDetail
	err := db.connection.Where("service_id = ?", serviceID).Where("group_id = ?", groupID).Take(&serviceDetail).Error
	return serviceDetail, err
}

func (db *reservationConnection) FindCageDetailByID(cageDetailID string) (entity.CageDetail, error) {
	var cageDetail entity.CageDetail
	err := db.connection.Where("id_cage_detail = ?", cageDetailID).Take(&cageDetail).Error
	return cageDetail, err
}

func (db *reservationConnection) FindProductByID(productID string) (entity.Product, error) {
	var product entity.Product
	err := db.connection.Where("id_product = ?", productID).Take(&product).Error
	return product, err
}
func (db *reservationConnection) FindGroupByID(groupID string) (entity.Group, error) {
	var group entity.Group
	err := db.connection.Where("id_group = ?", groupID).Take(&group).Error
	return group, err
}

func (db *reservationConnection) DeleteReservation(reservation entity.Reservation) error {
	var err error
	var reservationDetails []entity.ReservationDetail
	var reservationService entity.ReservationService
	var reservationInventory entity.ReservationInventory
	err = db.connection.Where("reservation_id = ?", reservation.IDReservation).Find(&reservationDetails).Error
	for _, reservationDetail := range reservationDetails {
		db.connection.Where("reservation_detail_id = ?", reservationDetail.IDReservationDetail).Delete(&reservationService)
		db.connection.Where("reservation_detail_id = ?", reservationDetail.IDReservationDetail).Delete(&reservationInventory)
	}
	err = db.connection.Where("reservation_id = ?", reservation.IDReservation).Delete(&reservationDetails).Error
	err = db.connection.Where("id_reservation = ?", reservation.IDReservation).Delete(&reservation).Error
	return err
}

func (db *reservationConnection) GetAllGroupDetail(hotelID string) ([]entity.GroupDetail, error) {
	var groupDetails []entity.GroupDetail
	err := db.connection.Where("hotel_id = ?", hotelID).Find(&groupDetails).Error
	return groupDetails, err
}

func (db *reservationConnection) GetAllReservation(hotelID string, filterPagination dto.FilterPagination) ([]entity.Reservation, dto.Pagination, error) {
	search := filterPagination.Search
	sortBy := filterPagination.SortBy
	orderBy := filterPagination.OrderBy
	perPage := int(filterPagination.PerPage)
	page := int(filterPagination.Page)
	fmt.Println(hotelID)

	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	var total int64

	var reservations []entity.Reservation
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(reservations.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
		}
	}

	listSortBy := []string{"name"}
	listSortOrder := []string{"desc", "asc"}

	if sortBy != "" && contains(listSortBy, sortBy) == true && orderBy != "" && contains(listSortOrder, orderBy) {
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	} else {
		sortBy = "created_at"
		orderBy = "desc"
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	}

	err := query.Where("hotel_id = ?", hotelID).Limit(perPage).Offset((page - 1) * perPage).
		Preload("Hotel").
		Preload("User").
		Preload("ReservationDetails").
		Preload("ReservationDetails.Reservation").
		Preload("Rate").
		Find(&reservations).
		Count(&total).
		Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return reservations, pagination, err
}

//func (db *speciesConnection) CreateReservation(reservation entity.Reservation) (entity.Reservation, error) {
//	err := db.connection.Save(&reservation).Error
//	db.connection.Find(&reservation)
//	return reservation, err
//}
//
//func (db *speciesConnection) CreateReservationDetail(reservationDetail entity.ReservationDetail) (entity.ReservationDetail, error) {
//	err := db.connection.Save(&reservationDetail).Error
//	db.connection.Find(&reservationDetail)
//	return reservationDetail, err
//}
//
//func (db *speciesConnection) CreateReservationService(reservationService entity.ReservationService) (entity.ReservationService, error) {
//	err := db.connection.Save(&reservationService).Error
//	db.connection.Find(&reservationService)
//	return reservationService, err
//}
//
//func (db *speciesConnection) CreateReservationInventory(reservationInventory entity.ReservationInventory) (entity.ReservationInventory, error) {
//	err := db.connection.Save(&reservationInventory).Error
//	db.connection.Find(&reservationInventory)
//	return reservationInventory, err
//}

// InsertReservation is to add reservation in database
func (db *reservationConnection) InsertReservation(reservation entity.Reservation, reservationDetails []entity.ReservationDetail, reservationServices []entity.ReservationService, reservationInventories []entity.ReservationInventory, transactionDetails []dto.TransactionDetail) (entity.Reservation, error) {
	err := db.connection.Save(&reservation).Error

	for _, transDetail := range transactionDetails {
		transDetail.ReservationDetail.ReservationID = reservation.IDReservation
		err = db.connection.Save(&transDetail.ReservationDetail).Error
		for _, transService := range transDetail.ReservationServices {
			transService.ReservationDetailID = transDetail.ReservationDetail.IDReservationDetail
			err = db.connection.Save(&transService).Error
		}
		for _, transProduct := range transDetail.ReservationProducts {
			transProduct.ReservationDetailID = transDetail.ReservationDetail.IDReservationDetail
			err = db.connection.Save(&transProduct).Error
		}
		for _, transInventory := range transDetail.ReservationInventories {
			transInventory.ReservationDetailID = transDetail.ReservationDetail.IDReservationDetail
			err = db.connection.Save(&transInventory).Error
		}
	}

	//for _, reservationDetail := range reservationDetails {
	//	reservationDetail.ReservationID = reservation.IDReservation
	//	//fmt.Println("Iki id detail " + reservationDetail.ReservationID)
	//	err = db.connection.Save(&reservationDetail).Error
	//	//fmt.Println(reservationDetail.IDReservationDetail)
	//	for _, reservationService := range reservationServices {
	//		reservationService.ReservationDetailID = reservationDetail.IDReservationDetail
	//		err = db.connection.Save(&reservationService).Error
	//	}
	//	for _, reservationInventory := range reservationInventories {
	//		reservationInventory.ReservationDetailID = reservationDetail.IDReservationDetail
	//		err = db.connection.Save(&reservationInventory).Error
	//	}
	//}

	db.connection.Preload("Hotel").Preload("User").Preload("ReservationDetails").Preload("ReservationDetails.Reservation").Find(&reservation)

	return reservation, err
}

// UpdateReservation is func to edit reservation in database
func (db *reservationConnection) UpdateReservation(reservation entity.Reservation, reservationDetails []entity.ReservationDetail, reservationServices []entity.ReservationService, reservationInventories []entity.ReservationInventory) (entity.Reservation, error) {
	var err error
	for _, reservationInventory := range reservationInventories {
		err = db.connection.Where("id_reservation_inventory = ?", reservationInventory.IDReservationInventory).Updates(&reservationInventory).Error
		if err != nil {
			return entity.Reservation{}, err
		}
	}
	fmt.Println(reservationServices)
	for _, reservationService := range reservationServices {
		err = db.connection.Where("id_reservation_service = ?", reservationService.IDReservationService).Updates(&reservationService).Error
		if err != nil {
			return entity.Reservation{}, err
		}
	}
	for _, reservationDetail := range reservationDetails {
		err = db.connection.Where("id_reservation_detail = ?", reservationDetail.IDReservationDetail).Updates(&reservationDetail).Error
		if err != nil {
			return entity.Reservation{}, err
		}
	}
	err = db.connection.Where("id_reservation = ?", reservation.IDReservation).Updates(&reservation).Error
	db.connection.Where("id_reservation = ?", reservation.IDReservation).Find(&reservation)
	return reservation, err
}

// FindReservationByID is func to get reservation by email
func (db *reservationConnection) FindReservationByID(reservationID string) (entity.Reservation, error) {
	var reservation entity.Reservation
	err := db.connection.
		Joins("JOIN reservation_details ON reservations.id_reservation=reservation_details.reservation_id").
		//Joins("JOIN reservation_services ON reservation_details.id_reservation_detail=reservation_services.reservation_detail_id").
		//Joins("JOIN reservation_inventories ON reservation_details.id_reservation_detail=reservation_inventories.reservation_detail_id").
		Where("id_reservation = ?", reservationID).
		Preload("Hotel").
		Preload("User").
		Preload("ReservationDetails.ReservationServices").
		Preload("ReservationDetails.ReservationInventories").
		//Preload("ReservationDetails.Reservation").
		//Preload("ReservationServices").
		//Preload("ReservationServices.ReservationDetail").
		//Preload("ReservationInventories").
		//Preload("ReservationInventories.ReservationDetail").
		Take(&reservation).Error
	return reservation, err
}

// NewReservationRepository is creates a new instance of ReservationRepository
func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &reservationConnection{
		connection: db,
	}
}
