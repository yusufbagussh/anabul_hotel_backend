package service

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/config"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
	"math"
	"time"
)

type TransactionDetail struct {
	ReservationDetail      entity.ReservationDetail
	ReservationServices    []entity.ReservationService
	ReservationInventories []entity.ReservationInventory
}

//type transactionData struct {
//	reservationDetails []transactionDetail
//}

// ReservationService is a contract of what reservationService can do
type ReservationService interface {
	GetAllReservation(hotelID string, filterPage dto.FilterPagination) ([]entity.Reservation, dto.Pagination, error)
	CreateReservation(reservation dto.CreateReservation, userID string) (interface{}, error)
	UpdateReservation(reservation dto.UpdateReservation) (entity.Reservation, error)
	DeleteReservation(reservationID string, userHotelID string) (error, interface{})
	ShowReservation(reservationID string, userHotelID string) (entity.Reservation, error, interface{})
	UpdatePaymentStatus(paymentStatus dto.UpdatePaymentStatus) (entity.Reservation, error)
	UpdateReservationStatus(paymentStatus dto.UpdateReservationStatus) (entity.Reservation, error)
	UpdateCheckInStatus(paymentStatus dto.UpdateCheckInStatus) (entity.Reservation, error)
}

type reservationService struct {
	reservationRepository repository.ReservationRepository
}

// NewReservationService creates a new instance of ReservationService
func NewReservationService(reservationRepo repository.ReservationRepository) ReservationService {
	return &reservationService{
		reservationRepository: reservationRepo,
	}
}

func (u *reservationService) UpdatePaymentStatus(paymentStatus dto.UpdatePaymentStatus) (entity.Reservation, error) {
	updatedReservation, err := u.reservationRepository.UpdatePaymentStatus(paymentStatus)
	if err != nil {
		return entity.Reservation{}, err
	}

	var createNotification entity.ReservationNotification
	createNotification.ReservationID = updatedReservation.IDReservation
	createNotification.Title = "Notifikasi Pembayaran"
	createNotification.Description = "Terimakasih anda melakukan pembayaran reservasi pet hotel blabla"
	createNotification.UserID = updatedReservation.UserID

	notificationData, errCreate := u.reservationRepository.CreateNotification(createNotification)
	if errCreate != nil {
		return entity.Reservation{}, err
	}

	app, _, _ := config.SetupFirebase()
	errSend := sendNotificationReservationStatus(app, notificationData)

	if errSend != nil {
		return entity.Reservation{}, errSend
	}

	return updatedReservation, err
}

func (u *reservationService) UpdateReservationStatus(reservationStatus dto.UpdateReservationStatus) (entity.Reservation, error) {
	updatedReservation, err := u.reservationRepository.UpdateReservationStatus(reservationStatus)
	if err != nil {
		return entity.Reservation{}, err
	}

	var createNotification entity.ReservationNotification
	createNotification.ReservationID = updatedReservation.IDReservation
	createNotification.Title = "Notifikasi Reservasi"
	if updatedReservation.ReservationStatus == "Proses" {
		createNotification.Description = "Reservasi anda sedang diproses dan akan dilakukan validasi oleh pihal admin, silahkan mohon tunggu sebentar notifikasi dari kami"
	} else if updatedReservation.ReservationStatus == "Diterima" {
		createNotification.Description = "Reservasi anda sudah kami terima, anda sudah bisa membawa hewan peliharaan anda ke tempat penitipan kami"
	} else if updatedReservation.ReservationStatus == "Ditolak" {
		createNotification.Description = "Mohon maaf reservasi anda kami tolak, karena tidak memenuhi syarat dan peraturan dari tempat penitipan kami"
	} else if updatedReservation.ReservationStatus == "Dibatalkan" {
		createNotification.Description = "Anda telah membatalkan layanan reservasi."
	}
	createNotification.UserID = updatedReservation.UserID

	notificationData, errCreate := u.reservationRepository.CreateNotification(createNotification)
	if errCreate != nil {
		return entity.Reservation{}, err
	}

	app, _, _ := config.SetupFirebase()
	errSend := sendNotificationReservationStatus(app, notificationData)

	if errSend != nil {
		return entity.Reservation{}, errSend
	}

	return updatedReservation, err
}

func (u *reservationService) UpdateCheckInStatus(checkInStatus dto.UpdateCheckInStatus) (entity.Reservation, error) {
	updatedReservation, err := u.reservationRepository.UpdateCheckInStatus(checkInStatus)
	if err != nil {
		return entity.Reservation{}, err
	}

	var createNotification entity.ReservationNotification
	createNotification.ReservationID = updatedReservation.IDReservation
	createNotification.Title = "Notifikasi Penitipan"
	if updatedReservation.CheckInStatus == "Masuk" {
		createNotification.Description = "Hewan peliharaan kesayangan Anda sudah masuk ke tempat penitipan hewan"
	} else {
		createNotification.Description = "Hewan peliharaan kesayangan Anda sudah keluar dari tempat penitipan hewan. Terima kasih sudah menggunakan layanan dari kami"
	}
	createNotification.UserID = updatedReservation.UserID

	notificationData, errCreate := u.reservationRepository.CreateNotification(createNotification)
	if errCreate != nil {
		return entity.Reservation{}, err
	}

	app, _, _ := config.SetupFirebase()
	errSend := sendNotificationReservationStatus(app, notificationData)

	if errSend != nil {
		return entity.Reservation{}, errSend
	}

	return updatedReservation, err
}

func sendNotificationReservationStatus(app *firebase.App, notificationData entity.ReservationNotification) error {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		//log.Fatalf("error getting Messaging client: %v\n", err)
		return err
	}

	fmt.Println("Tes 1")

	registrationToken := "dVib4muX6vS4vk3SUDKpRG:APA91bGfNG2lQSubB6o5v9D0Aj5ClnhW3DEgPbZ0fifj-Yh0hJ_U40OsUGtp6Zd6jM_YDYwOu1Zx7ZTGKAKIvA9aoZU_YzZMq1gw_EoI-UyiPitpHHC4Z6Fl8hSFxNXAzqSm9QG1KQHz"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: notificationData.Title,
			Body:  notificationData.Description,
		},
		Token: registrationToken,
	}

	fmt.Println("Tes 2")

	response, errSend := client.Send(ctx, message)
	if errSend != nil {
		//log.Fatalln(err)
		return errSend
	}

	fmt.Println("Tes 3")

	fmt.Println("Successfully sent message:", response)
	return nil
}

func (u *reservationService) CreateReservation(reservation dto.CreateReservation, userID string) (interface{}, error) {
	reservationEntity := entity.Reservation{}
	var reservationDetailEntity entity.ReservationDetail
	var reservationServiceEntity entity.ReservationService
	var reservationInventoryEntity entity.ReservationInventory
	var reservationProductEntity entity.ReservationProduct
	var errMap error
	reservationEntity.HotelID = reservation.HotelID
	reservationEntity.UserID = userID
	reservationEntity.ReservationStatus = reservation.ReservationStatus
	reservationEntity.CheckInStatus = reservation.CheckInStatus
	reservationEntity.PaymentStatus = reservation.PaymentStatus
	//reservationEntity.CreatedBy = reservation.CreatedBy
	//reservationEntity.UpdatedBy = reservation.UpdatedBy
	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Jakarta")
	StartDate, _ := time.ParseInLocation(layout, reservation.StartDate, loc)
	EndDate, _ := time.ParseInLocation(layout, reservation.EndDate, loc)
	diff := EndDate.Sub(StartDate)
	//fmt.Println(math.Ceil(diff.Hours()/24) + 1) // number of days
	diffDate := math.Ceil(diff.Hours()/24) + 1
	reservationEntity.StartDate = StartDate
	reservationEntity.EndDate = EndDate
	reservationEntity.TotalCost = reservation.TotalCost
	reservationEntity.DPCost = reservation.DPCost

	var reservationDetails []entity.ReservationDetail
	var reservationServices []entity.ReservationService
	var reservationInventories []entity.ReservationInventory
	var reservationProducts []entity.ReservationProduct

	var cageCost float64
	var serviceCost float64
	//var productCost float64
	var totalProductCost float64
	var totalCageCost float64

	var reservDetail dto.TransactionDetail
	var reservDetails []dto.TransactionDetail
	for _, reservationDetail := range reservation.ReservationDetail {
		errMap = smapping.FillStruct(&reservationDetailEntity, smapping.MapFields(&reservationDetail))
		if errMap != nil {
			return reservationEntity, errMap
		}
		reservationDetails = append(reservationDetails, reservationDetailEntity)
		reservDetail.ReservationDetail = reservationDetailEntity
		cageDetail, _ := u.reservationRepository.FindCageDetailByID(reservationDetail.CageDetailID)
		//product, _ := u.reservationRepository.FindProductByID(reservationDetail.ProductID)
		cageCost += cageDetail.Price
		//productCost += product.Price * 2

		listGroupDetail, _ := u.reservationRepository.GetAllGroupDetail(reservation.HotelID)
		//fmt.Println(listGroupDetail)
		var groupID string
		for _, groupDetail := range listGroupDetail {
			if groupDetail.SpeciesID != "" && groupDetail.MinWeight == 0 && groupDetail.MaxWeight == 0 {
				if groupDetail.SpeciesID == reservationDetail.SpeciesID {
					groupID = groupDetail.GroupID
				}
			}
			if groupDetail.SpeciesID != "" && groupDetail.MinWeight != 0 && groupDetail.MaxWeight != 0 {
				if groupDetail.SpeciesID == reservationDetail.SpeciesID && reservationDetail.Weight >= groupDetail.MinWeight && reservationDetail.Weight <= groupDetail.MaxWeight {
					groupID = groupDetail.GroupID
				}

			}
		}

		var qtyStay uint8
		for _, reservationServiceDetail := range reservationDetail.ReservationService {
			reservDetail.ReservationServices = []entity.ReservationService{}
			errMap = smapping.FillStruct(&reservationServiceEntity, smapping.MapFields(&reservationServiceDetail))
			if errMap != nil {
				return reservationEntity, errMap
			}
			serviceDetail, _ := u.reservationRepository.FindServiceDetailByID(reservationServiceDetail.ServiceID, groupID)
			serviceCost += serviceDetail.Price * float64(reservationServiceDetail.Quantity)
			qtyStay += reservationServiceDetail.Quantity
			//totalProductCost += productCost * float64(reservationServiceDetail.Quantity)
			reservationServices = append(reservationServices, reservationServiceEntity)
			reservDetail.ReservationServices = append(reservDetail.ReservationServices, reservationServiceEntity)
		}
		//totalProductCost = productCost * diffDate
		totalCageCost = cageCost * diffDate

		for _, reservationProduct := range reservationDetail.ReservationProduct {
			errMap = smapping.FillStruct(&reservationProductEntity, smapping.MapFields(&reservationProduct))
			if errMap != nil {
				return reservationEntity, errMap
			}
			productData, _ := u.reservationRepository.FindProductByID(reservationProduct.ProductID)
			totalProductCost += productData.Price * float64(reservationProduct.Quantity)
			reservationProducts = append(reservationProducts, reservationProductEntity)
			reservDetail.ReservationProducts = append(reservDetail.ReservationProducts, reservationProductEntity)
		}

		for _, reservationInventory := range reservationDetail.ReservationInventory {
			reservDetail.ReservationInventories = []entity.ReservationInventory{}
			errMap = smapping.FillStruct(&reservationInventoryEntity, smapping.MapFields(&reservationInventory))
			if errMap != nil {
				return reservationEntity, errMap
			}
			reservationInventories = append(reservationInventories, reservationInventoryEntity)
			reservDetail.ReservationInventories = append(reservDetail.ReservationInventories, reservationInventoryEntity)
		}
		reservDetails = append(reservDetails, reservDetail)
	}
	fmt.Println(reservDetails)
	fmt.Println(totalProductCost)
	fmt.Println(totalCageCost)
	fmt.Println(serviceCost)
	reservationEntity.TotalCost = totalProductCost + totalCageCost + serviceCost

	//errMap := smapping.FillStruct(&reservationEntity, smapping.MapFields(&reservation))
	//if errMap != nil {
	//	return reservationEntity, errMap
	//}
	updatedReservation, err := u.reservationRepository.InsertReservation(reservationEntity, reservationDetails, reservationServices, reservationInventories, reservDetails)
	return updatedReservation, err
}
func (u *reservationService) UpdateReservation(reservation dto.UpdateReservation) (entity.Reservation, error) {
	reservationEntity := entity.Reservation{}
	var reservationDetailEntity entity.ReservationDetail
	var reservationServiceEntity entity.ReservationService
	var reservationInventoryEntity entity.ReservationInventory
	var reservationProductEntity entity.ReservationProduct
	var errMap error
	reservationEntity.IDReservation = reservation.IDReservation
	reservationEntity.HotelID = reservation.HotelID
	reservationEntity.UserID = reservation.UserID
	reservationEntity.ReservationStatus = reservation.ReservationStatus
	reservationEntity.CheckInStatus = reservation.CheckInStatus
	reservationEntity.PaymentStatus = reservation.PaymentStatus
	//reservationEntity.CreatedBy = reservation.CreatedBy
	//reservationEntity.UpdatedBy = reservation.UpdatedBy
	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Jakarta")
	StartDate, _ := time.ParseInLocation(layout, reservation.StartDate, loc)
	EndDate, _ := time.ParseInLocation(layout, reservation.EndDate, loc)
	diff := EndDate.Sub(StartDate)
	//fmt.Println(math.Ceil(diff.Hours()/24) + 1) // number of days
	diffDate := math.Ceil(diff.Hours()/24) + 1
	reservationEntity.StartDate = StartDate
	reservationEntity.EndDate = EndDate
	reservationEntity.TotalCost = reservation.TotalCost
	reservationEntity.DPCost = reservation.DPCost

	var reservationDetails []entity.ReservationDetail
	var reservationServices []entity.ReservationService
	var reservationInventories []entity.ReservationInventory
	var reservationProducts []entity.ReservationProduct

	var cageCost float64
	var serviceCost float64
	//var productCost float64
	var totalProductCost float64
	var totalCageCost float64

	//var totalCostPerPet float64

	for _, reservationDetail := range reservation.ReservationDetail {
		errMap = smapping.FillStruct(&reservationDetailEntity, smapping.MapFields(&reservationDetail))
		if errMap != nil {
			return reservationEntity, errMap
		}
		reservationDetails = append(reservationDetails, reservationDetailEntity)
		cageDetail, _ := u.reservationRepository.FindCageDetailByID(reservationDetail.CageDetailID)
		//product, _ := u.reservationRepository.FindProductByID(reservationDetail.ProductID)
		cageCost += cageDetail.Price
		//productCost += product.Price * 2

		listGroupDetail, _ := u.reservationRepository.GetAllGroupDetail(reservation.HotelID)
		var groupID string
		for _, groupDetail := range listGroupDetail {
			if groupDetail.SpeciesID != "" && groupDetail.MinWeight == 0 && groupDetail.MaxWeight == 0 {
				if groupDetail.SpeciesID == reservationDetail.SpeciesID {
					groupID = groupDetail.GroupID
				}
			}
			if groupDetail.SpeciesID != "" && groupDetail.MinWeight != 0 && groupDetail.MaxWeight != 0 {
				if groupDetail.SpeciesID == reservationDetail.SpeciesID && reservationDetail.Weight >= groupDetail.MinWeight && reservationDetail.Weight <= groupDetail.MaxWeight {
					groupID = groupDetail.GroupID
				}

			}
		}

		var qtyStay uint8
		for _, reservationServiceDetail := range reservationDetail.ReservationService {
			errMap = smapping.FillStruct(&reservationServiceEntity, smapping.MapFields(&reservationServiceDetail))
			if errMap != nil {
				return reservationEntity, errMap
			}
			serviceDetail, _ := u.reservationRepository.FindServiceDetailByID(reservationServiceDetail.ServiceID, groupID)
			serviceCost += serviceDetail.Price * float64(reservationServiceDetail.Quantity)
			qtyStay += reservationServiceDetail.Quantity
			//totalProductCost += productCost * float64(reservationServiceDetail.Quantity)
			reservationServices = append(reservationServices, reservationServiceEntity)
		}
		//totalProductCost = productCost * diffDate
		totalCageCost = cageCost * diffDate

		for _, reservationProduct := range reservationDetail.ReservationProduct {
			errMap = smapping.FillStruct(&reservationProductEntity, smapping.MapFields(&reservationProduct))
			if errMap != nil {
				return reservationEntity, errMap
			}
			reservationProducts = append(reservationProducts, reservationProductEntity)
			productData, _ := u.reservationRepository.FindProductByID(reservationProduct.ProductID)
			totalProductCost += productData.Price * float64(reservationProduct.Quantity)

		}

		for _, reservationInventory := range reservationDetail.ReservationInventory {
			errMap = smapping.FillStruct(&reservationInventoryEntity, smapping.MapFields(&reservationInventory))
			if errMap != nil {
				return reservationEntity, errMap
			}
			reservationInventories = append(reservationInventories, reservationInventoryEntity)
		}
	}
	fmt.Println(totalProductCost)
	fmt.Println(totalCageCost)
	fmt.Println(serviceCost)
	reservationEntity.TotalCost = totalProductCost + totalCageCost + serviceCost

	//errMap := smapping.FillStruct(&reservationEntity, smapping.MapFields(&reservation))
	//if errMap != nil {
	//	return reservationEntity, errMap
	//}
	updatedReservation, err := u.reservationRepository.UpdateReservation(reservationEntity, reservationDetails, reservationServices, reservationInventories)
	return updatedReservation, err
}
func (u *reservationService) DeleteReservation(reservationID string, userHotelID string) (error, interface{}) {
	reservation, err := u.reservationRepository.FindReservationByID(reservationID)
	if reservation.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	errDel := u.reservationRepository.DeleteReservation(reservation)
	return errDel, nil
}
func (u *reservationService) ShowReservation(reservationID string, userHotelID string) (entity.Reservation, error, interface{}) {
	result, err := u.reservationRepository.FindReservationByID(reservationID)
	if result.HotelID != userHotelID {
		return entity.Reservation{}, nil, false
	}
	return result, err, nil
}
func (u *reservationService) GetAllReservation(hotelID string, filterPage dto.FilterPagination) ([]entity.Reservation, dto.Pagination, error) {
	admins, pagination, err := u.reservationRepository.GetAllReservation(hotelID, filterPage)
	return admins, pagination, err
}
