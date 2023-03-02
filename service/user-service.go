package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
	"os"
	"path/filepath"
)

// UserService is a contract of what userService can do
type UserService interface {
	Update(user dto.UserHotelUpdateDTO, ctx *gin.Context) (entity.User, error)
	GetProfile(userID string) (entity.User, error)
	ChangePassword(newDataPass dto.ChangePasswordDTO, userID string) (interface{}, error)
	AllAdmin(filterPagination dto.FilterPagination) ([]entity.User, dto.Pagination, error)
	AllStaff(hotelID string, pagination dto.FilterPagination) ([]entity.User, dto.Pagination, error)
	CreateUser(user dto.UserHotelCreateDTO, ctx *gin.Context) (entity.User, error)
	UpdateUser(user dto.UserHotelUpdateDTO, ctx *gin.Context) (entity.User, error)
	DeleteUser(userID string) error
	ShowUser(userID string) (entity.User, error)
	IsAllowedUser(userID string, hotelID string) bool
	DeleteStaff(userID string, userHotelID string) (error, interface{})
	ShowStaff(userID string, userHotelID string) (entity.User, error, interface{})
	SaveDeviceToken(device dto.Device, userID string) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (u *userService) SaveDeviceToken(device dto.Device, userID string) (entity.User, error) {
	user, errFind := u.userRepository.FindUserByID(userID)
	if errFind != nil {
		return entity.User{}, errFind
	}
	user.DeviceToken = device.DeviceToken
	userUpdate, errSave := u.userRepository.SaveDeviceToken(user)
	if errSave != nil {
		return entity.User{}, errSave
	}
	return userUpdate, nil
}

func (u *userService) CreateUser(user dto.UserHotelCreateDTO, ctx *gin.Context) (entity.User, error) {
	userCreate := entity.User{}

	if user.Image != nil {
		extension := filepath.Ext(user.Image.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.New().String() + extension
		userCreate.Image = newImage

		err := ctx.SaveUploadedFile(user.Image, "uploads/user-profiles/"+newImage)
		if err != nil {
			return userCreate, err
		}
	}

	if user.KTP != nil {
		fileName := user.KTP.Filename
		extension := filepath.Ext(user.KTP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newKTPFile := fileName + extension
		userCreate.KTP = newKTPFile

		err := ctx.SaveUploadedFile(user.KTP, "uploads/user-profiles/"+newKTPFile)
		if err != nil {
			return userCreate, err
		}
	}

	userCreate.Name = user.UserName
	userCreate.Email = user.UserEmail
	userCreate.Password = user.Password
	userCreate.Address = user.Address
	userCreate.Phone = user.Phone
	userCreate.Role = user.Role
	userCreate.HotelID = user.HotelID
	userCreate.NIK = user.NIK
	updatedUser, err := u.userRepository.InsertUser(userCreate)
	return updatedUser, err
}
func (u *userService) UpdateUser(user dto.UserHotelUpdateDTO, ctx *gin.Context) (entity.User, error) {
	userUpdate := entity.User{}

	if user.Image != nil {
		extension := filepath.Ext(user.Image.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.New().String() + extension
		userUpdate.Image = newImage

		err := ctx.SaveUploadedFile(user.Image, "uploads/user-profiles/"+newImage)
		if err != nil {
			return userUpdate, err
		} else {
			userByID, _ := u.userRepository.FindUserByID(user.ID)
			if userByID.Image != "" {
				errRemove := os.Remove("./uploads/user-profiles/" + userByID.Image)
				if errRemove != nil {
					return userUpdate, errRemove
				}
			}
		}
	}

	if user.KTP != nil {
		fileName := user.KTP.Filename
		extension := filepath.Ext(user.KTP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newKTPFile := fileName + extension
		userUpdate.KTP = newKTPFile

		err := ctx.SaveUploadedFile(user.KTP, "uploads/user-profiles/"+newKTPFile)
		if err != nil {
			return userUpdate, err
		} else {
			userByID, _ := u.userRepository.FindUserByID(user.ID)
			if userByID.KTP != "" {
				errRemove := os.Remove("./uploads/user-profiles/" + userByID.KTP)
				if errRemove != nil {
					return userUpdate, errRemove
				}
			}
		}
	}
	userUpdate.ID = user.ID
	userUpdate.Name = user.Name
	userUpdate.Email = user.Email
	userUpdate.Address = user.Address
	userUpdate.Phone = user.Phone
	userUpdate.NIK = user.NIK
	userUpdate.Role = user.Role
	userUpdate.HotelID = user.HotelID
	updatedUser, err := u.userRepository.UpdateUser(userUpdate)
	return updatedUser, err
}
func (u *userService) DeleteUser(userID string) error {
	user, err := u.userRepository.FindUserByID(userID)
	if err != nil {
		return err
	}
	if user.Image != "" {
		errRemove := os.Remove("./uploads/user-profiles/" + user.Image)
		if errRemove != nil {
			return errRemove
		}
	}
	if user.KTP != "" {
		errRemove := os.Remove("./uploads/user-profiles/" + user.KTP)
		if errRemove != nil {
			return errRemove
		}
	}
	errDel := u.userRepository.DeleteUser(user)
	return errDel
}
func (u *userService) DeleteStaff(userID string, userHotelID string) (error, interface{}) {
	user, err := u.userRepository.FindUserByID(userID)
	if user.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	if user.Image != "" {
		errRemove := os.Remove("./uploads/user-profiles/" + user.Image)
		if errRemove != nil {
			return errRemove, nil
		}
	}
	if user.KTP != "" {
		errRemove := os.Remove("./uploads/user-profiles/" + user.KTP)
		if errRemove != nil {
			return errRemove, nil
		}
	}
	errDel := u.userRepository.DeleteUser(user)
	return errDel, nil
}
func (u *userService) ShowStaff(userID string, userHotelID string) (entity.User, error, interface{}) {
	result, err := u.userRepository.FindUserByID(userID)
	if result.HotelID != userHotelID {
		return entity.User{}, nil, false
	}
	return result, err, nil
}
func (u *userService) ShowUser(userID string) (entity.User, error) {
	result, err := u.userRepository.FindUserByID(userID)
	return result, err
}
func (u *userService) AllAdmin(filterPagination dto.FilterPagination) ([]entity.User, dto.Pagination, error) {
	admins, pagination, err := u.userRepository.AllAdmin(filterPagination)
	return admins, pagination, err
}
func (u *userService) AllStaff(hotelID string, filterPagination dto.FilterPagination) ([]entity.User, dto.Pagination, error) {
	staffs, pagination, err := u.userRepository.AllStaff(hotelID, filterPagination)
	return staffs, pagination, err
}

func (u *userService) ChangePassword(newDataPass dto.ChangePasswordDTO, userID string) (interface{}, error) {
	//res, err := u.userRepository.FindByEmail(newDataPass.Email)
	oldDataPass, err := u.userRepository.FindUserByID(userID)
	if err != nil {
		return nil, err
	} else {
		if comparePassword(oldDataPass.Password, []byte(newDataPass.OldPassword)) {
			oldDataPass.Password = newDataPass.NewPassword
		}
		result, err := u.userRepository.ChangePass(oldDataPass)
		return result, err
	}
}

func (u *userService) Update(user dto.UserHotelUpdateDTO, ctx *gin.Context) (entity.User, error) {
	userUpdate := entity.User{}

	if user.Image != nil {
		//get extention of file image
		extension := filepath.Ext(user.Image.Filename)
		// Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.New().String() + extension
		userUpdate.Image = newImage
		//save image to uploads/images
		err := ctx.SaveUploadedFile(user.Image, "uploads/user-profiles/"+newImage)
		if err != nil {
			//ctx.String(http.StatusInternalServerError, "unknown error")
			return userUpdate, err
		} else {
			// Removing file from the directory
			//Using Remove() function
			userByID, _ := u.userRepository.FindUserByID(user.ID)
			if userByID.Image != "" {
				errRemove := os.Remove("./uploads/user-profiles/" + userByID.Image)
				if errRemove != nil {
					return userUpdate, errRemove
				}
			}
		}
	}
	if user.KTP != nil {
		fileName := user.KTP.Filename
		extension := filepath.Ext(user.KTP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newNIKFile := fileName + extension
		userUpdate.KTP = newNIKFile

		err := ctx.SaveUploadedFile(user.KTP, "uploads/user-profiles/"+newNIKFile)
		if err != nil {
			return userUpdate, err
		} else {
			userByID, _ := u.userRepository.FindUserByID(user.ID)
			if userByID.KTP != "" {
				errRemove := os.Remove("./uploads/user-profiles/" + userByID.KTP)
				if errRemove != nil {
					return userUpdate, errRemove
				}
			}
		}
	}
	userUpdate.ID = user.ID
	userUpdate.Name = user.Name
	userUpdate.Email = user.Email
	userUpdate.Address = user.Address
	userUpdate.Phone = user.Phone
	userUpdate.Role = user.Role
	userUpdate.HotelID = user.HotelID
	userUpdate.NIK = user.NIK
	updatedUser, err := u.userRepository.UpdateUser(userUpdate)
	return updatedUser, err
}

func (u *userService) GetProfile(userID string) (entity.User, error) {
	user, err := u.userRepository.FindUserByID(userID)
	return user, err
}

func (u *userService) IsAllowedUser(hotelID string, userHotelID string) bool {
	//userHotel, _ := u.userRepository.FindUserByID(userID)
	//if userHotel.HotelID == hotelID {
	//	return true
	//} else {
	//	return true
	//}
	if hotelID == userHotelID {
		return true
	} else {
		return false
	}
}
