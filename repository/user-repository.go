package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"math"
	"strings"
)

// UserRepository is contract what userRepository can do to db
type UserRepository interface {
	AllAdmin(filterPagination dto.FilterPagination) ([]entity.User, dto.Pagination, error)
	AllStaff(hotelID string, filterPagination dto.FilterPagination) ([]entity.User, dto.Pagination, error)
	InsertUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	VerifyCredential(email string) interface{}
	FindByEmail(email string) (entity.User, error)
	FindUserByID(userID string) (entity.User, error)
	ChangePass(user entity.User) (entity.User, error)
	UpdateVerified(user entity.User) (entity.User, error)
	SaveDeviceToken(user entity.User) (entity.User, error)
	DeleteUser(user entity.User) error
	//ProfileUser(userID string) entity.User
	//IsDuplicateEmail(email string) (tx *gorm.DB)
	//CheckPass(user entity.User, passwordResetToken string) (entity.User, interface{})
}

// userConnection adalah func untuk melakukan query data ke tabel user
type userConnection struct {
	connection *gorm.DB
}

// NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) DeleteUser(user entity.User) error {
	err := db.connection.Where("id = ?", user.ID).Delete(&user).Error
	return err
}

func (db *userConnection) AllAdmin(filterPagination dto.FilterPagination) ([]entity.User, dto.Pagination, error) {
	search := filterPagination.Search
	sortBy := filterPagination.SortBy
	orderBy := filterPagination.OrderBy
	perPage := int(filterPagination.PerPage)
	page := int(filterPagination.Page)

	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	var admins []entity.User
	query := db.connection.Joins("JOIN hotels ON user.hotel_id=hotels.id_hotel")

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(user.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
				Or("LOWER(hotels.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
		}
	}

	listSortBy := []string{"name", "email"}
	listSortOrder := []string{"desc", "asc"}

	if sortBy != "" && contains(listSortBy, sortBy) == true && orderBy != "" && contains(listSortOrder, orderBy) {
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	} else {
		sortBy = "created_at"
		orderBy = "desc"
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	}

	var total int64

	err := query.Where("role = ?", "Admin").Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&admins).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return admins, pagination, err
}

func (db *userConnection) AllStaff(hotelID string, filterPagination dto.FilterPagination) ([]entity.User, dto.Pagination, error) {
	search := filterPagination.Search
	sortBy := filterPagination.SortBy
	orderBy := filterPagination.OrderBy
	perPage := int(filterPagination.PerPage)
	page := int(filterPagination.Page)

	var staffs []entity.User
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(users.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
		}
	}

	listSortBy := []string{"name", "email"}
	listSortOrder := []string{"desc", "asc"}

	if sortBy != "" && contains(listSortBy, sortBy) == true && orderBy != "" && contains(listSortOrder, orderBy) {
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	} else {
		sortBy = "created_at"
		orderBy = "desc"
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	}

	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	var total int64

	err := query.Where("role = ?", "Staff").Where("hotel_id = ?", hotelID).Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&staffs).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return staffs, pagination, err
}

func (db *userConnection) ChangePass(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	//err := db.connection.Model(&user).Where("id = ?", user.ID).Updates(&entity.User{Password: user.Password}).Error
	err := db.connection.Where("id = ?", user.ID).Updates(&user).Error
	fmt.Println(user.HotelID)
	db.connection.Find(&user)
	return user, err
}

func (db *userConnection) UpdateVerified(user entity.User) (entity.User, error) {
	//err := db.connection.Model(&user).Where("id = ?", user.ID).Updates(&entity.User{Password: user.Password}).Error
	err := db.connection.Where("id = ?", user.ID).Updates(&user).Error
	fmt.Println(user.HotelID)
	db.connection.Find(&user)
	return user, err
}

func (db *userConnection) SaveDeviceToken(user entity.User) (entity.User, error) {
	err := db.connection.Model(&user).Where("id = ?", user.ID).Update("device_token", user.Password).Error
	db.connection.Find(&user)
	return user, err
}

// InsertUser is to add user in database
func (db *userConnection) InsertUser(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	err := db.connection.Save(&user).Error
	db.connection.Find(&user)
	return user, err
}

// UpdateUser is func to edit user in database
func (db *userConnection) UpdateUser(user entity.User) (entity.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		err := db.connection.Where("id = ?", user.ID).Take(&tempUser).Error
		if err != nil {
			return tempUser, err
		}
		user.Password = tempUser.Password
	}
	err := db.connection.Where("id = ?", user.ID).Updates(&user).Error
	db.connection.Where("id = ?", user.ID).Find(&user)
	return user, err
}

// VerifyCredential adalah func untuk melakukan pengecekan user berdasarkan email dan password
func (db *userConnection) VerifyCredential(email string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Preload("Hotel").Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

// IsDuplicateEmail adalah func untuk melakukan pengecekan apakah ada email yang sama
//func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
//	var user entity.User
//	return db.connection.Where("email = ?", email).Take(&user)
//}

// FindByEmail is func to get user by email
func (db *userConnection) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := db.connection.Where("email = ?", email).Preload("Hotel").Take(&user).Error
	return user, err
}

// FindUserByID is func to get user by email
func (db *userConnection) FindUserByID(userID string) (entity.User, error) {
	var user entity.User
	err := db.connection.Where("id = ?", userID).Preload("Hotel").Find(&user).Error
	return user, err
}

// ProfileUser is func to get user in book by book_user_id
//func (db *userConnection) ProfileUser(userID string) entity.User {
//	var user entity.User
//	db.connection.Find(&user, userID)
//	return user
//}

// hashAndSalt adalah function untuk memlakukan generate password secara bycript
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

//Tidak Terpakai
//func (db *userConnection) CheckPass(user entity.User, passwordResetToken string) (entity.User, interface{}) {
//	res := db.connection.First(&user, "password_reset_token = ? AND password_reset_at > ?", passwordResetToken, time.Now())
//	return user, res.Error
//}
