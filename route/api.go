package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/controller"
	"github.com/yusufbagussh/pet_hotel_backend/middleware"
	"github.com/yusufbagussh/pet_hotel_backend/service"
)

type Route interface {
	Routes(route *gin.Engine)
}

type route struct {
	userController                 controller.UserController
	authController                 controller.AuthController
	requestController              controller.RequestController
	hotelController                controller.HotelController
	provinceController             controller.ProvinceController
	cityController                 controller.CityController
	districtController             controller.DistrictController
	classController                controller.ClassController
	categoryController             controller.CategoryController
	speciesController              controller.SpeciesController
	petController                  controller.PetController
	groupController                controller.GroupController
	groupDetailController          controller.GroupDetailController
	cageCategoryController         controller.CageCategoryController
	cageTypeController             controller.CageTypeController
	cageDetailController           controller.CageDetailController
	cageController                 controller.CageController
	serviceController              controller.ServiceController
	serviceDetailController        controller.ServiceDetailController
	productController              controller.ProductController
	reservationController          controller.ReservationController
	reservationDetailController    controller.ReservationDetailController
	reservationInventoryController controller.ReservationInventoryController
	reservationServiceController   controller.ReservationServiceController
	reservationProductController   controller.ReservationProductController
	reservationConditionController controller.ReservationConditionController
	jwtService                     service.JWTService
	redisService                   service.RedisService
	userService                    service.UserService
	rateController                 controller.RateController
	responseController             controller.ResponseController
}

func (r *route) Routes(route *gin.Engine) {

	fmt.Println()
	//TODO implement me
	authUp := route.Group("api/auth")
	{
		authUp.POST("/login", r.authController.Login)
		authUp.POST("/register", r.authController.Register)
		authUp.POST("/forgotpass", r.authController.ForgotPass)
		authUp.PUT("/resetpass/:resetToken", r.authController.ResetPass)
		authUp.GET("/notif", r.authController.TestFCM)
		authUp.POST("/registerUser", r.authController.RegisterUser)
		authUp.PUT("/verify/:verificationCode", r.authController.ActivationEmail)
	}

	authIn := route.Group("api/auth", middleware.AuthorizeJWT(r.jwtService, r.userService, r.redisService))
	{
		authIn.POST("/refresh", r.authController.Refresh)
		authIn.POST("/logout", r.authController.Logout)
	}

	userRoutes := route.Group("api/user", middleware.AuthorizeJWT(r.jwtService, r.userService, r.redisService))
	{
		userRoutes.GET("/profile", r.userController.GetProfile)
		userRoutes.PUT("/profile", r.userController.UpdateProfile)
		userRoutes.PUT("/changepass", r.userController.ChangePassword)
	}

	province := route.Group("api/province")
	{
		province.GET("/all", r.provinceController.GetAllProvince)
		province.GET("/show/:id", r.provinceController.ShowProvince)
	}

	provinceRoutes := route.Group("api/province", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		provinceRoutes.POST("/add", r.provinceController.CreateProvince)
		provinceRoutes.PUT("/update", r.provinceController.UpdateProvince)
		provinceRoutes.DELETE("/delete/:id", r.provinceController.DeleteProvince)
	}

	city := route.Group("api/city")
	{
		city.GET("/all", r.cityController.GetAllCity)
		city.GET("/show/:id", r.cityController.ShowCity)
	}

	cityRoutes := route.Group("api/city", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		cityRoutes.POST("/add", r.cityController.CreateCity)
		cityRoutes.PUT("/update", r.cityController.UpdateCity)
		cityRoutes.DELETE("/delete/:id", r.cityController.DeleteCity)
	}

	district := route.Group("api/district")
	{
		district.GET("/all", r.districtController.GetAllDistrict)
		district.GET("/show/:id", r.districtController.ShowDistrict)
	}

	districtRoutes := route.Group("api/district", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		districtRoutes.POST("/add", r.districtController.CreateDistrict)
		districtRoutes.PUT("/update", r.districtController.UpdateDistrict)
		districtRoutes.DELETE("/delete/:id", r.districtController.DeleteDistrict)
	}

	request := route.Group("api/request")
	{
		request.POST("/add", r.requestController.CreateRequest)
	}

	requestRoutes := route.Group("api/request", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		requestRoutes.GET("/all", r.requestController.GetAllRequest)
		requestRoutes.GET("/show/:id", r.requestController.ShowRequest)
		requestRoutes.PUT("/status", r.requestController.UpdateRequestStatus)
		//requestRoutes.PUT("/update", r.requestController.UpdateRequest)
		requestRoutes.DELETE("/delete/:id", r.requestController.DeleteRequest)
	}

	hotel := route.Group("api/hotel", middleware.AuthorizeJWT(r.jwtService, r.userService, r.redisService))
	{
		hotel.GET("/all", r.hotelController.GetAllHotel)
		hotel.GET("/show/:id", r.hotelController.ShowHotel)
	}

	hotelRoutes := route.Group("api/hotel", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		//hotelRoutes.POST("/add", r.hotelController.CreateHotel)
		//hotelRoutes.POST("/admin", r.hotelController.CreateHotelAdmin)
		hotelRoutes.PUT("/update", r.hotelController.UpdateHotel)
		hotelRoutes.DELETE("/delete/:id", r.hotelController.DeleteHotel)
	}

	adminRoutes := route.Group("api/admin", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		adminRoutes.GET("/all", r.userController.GetAdmin)
		adminRoutes.GET("/show/:id", r.userController.ShowAdmin)
		//adminRoutes.POST("/add", r.userController.CreateUser)
		adminRoutes.PUT("/update", r.userController.UpdateAdmin)
		adminRoutes.DELETE("/delete/:id", r.userController.DeleteAdmin)
	}

	hotelAdminRoutes := route.Group("api/hotel", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		hotelAdminRoutes.GET("/profile", r.hotelController.GetProfileHotel)
		hotelAdminRoutes.PUT("/profile", r.hotelController.UpdateProfileHotel)
	}

	class := route.Group("api/class", middleware.AuthorizeJWT(r.jwtService, r.userService, r.redisService))
	{
		class.GET("/all", r.classController.GetAllClass)
		class.GET("/show/:id", r.classController.ShowClass)
	}

	classRoutes := route.Group("api/class", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		classRoutes.POST("/add", r.classController.CreateClass)
		classRoutes.PUT("/update", r.classController.UpdateClass)
		classRoutes.DELETE("/delete/:id", r.classController.DeleteClass)
	}

	category := route.Group("api/category")
	{
		category.GET("/all", r.categoryController.GetAllCategory)
		category.GET("/show/:id", r.categoryController.ShowCategory)
	}

	categoryRoutes := route.Group("api/category", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		categoryRoutes.POST("/add", r.categoryController.CreateCategory)
		categoryRoutes.PUT("/update", r.categoryController.UpdateCategory)
		categoryRoutes.DELETE("/delete/:id", r.categoryController.DeleteCategory)
	}

	species := route.Group("api/species")
	{
		species.GET("/all", r.speciesController.GetSpecies)
		species.GET("/show/:id", r.speciesController.ShowSpecies)
	}

	speciesRoutes := route.Group("api/species", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		speciesRoutes.POST("/add", r.speciesController.CreateSpecies)
		speciesRoutes.PUT("/update", r.speciesController.UpdateSpecies)
		speciesRoutes.DELETE("/delete/:id", r.speciesController.DeleteSpecies)
	}

	petRoutes := route.Group("api/pet", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Customer"))
	{
		petRoutes.GET("/all", r.petController.GetPet)
		petRoutes.GET("/show/:id", r.petController.ShowPet)
		petRoutes.POST("/add", r.petController.CreatePet)
		petRoutes.PUT("/update", r.petController.UpdatePet)
		petRoutes.DELETE("/delete/:id", r.petController.DeletePet)
	}

	staffRoutes := route.Group("api/staff", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		staffRoutes.GET("/all", r.userController.GetStaff)
		staffRoutes.GET("/show/:id", r.userController.ShowStaff)
		staffRoutes.POST("/add", r.userController.CreateUser)
		staffRoutes.PUT("/update", r.userController.UpdateStaff)
		staffRoutes.DELETE("/delete/:id", r.userController.DeleteStaff)
	}

	cage := route.Group("api/cage")
	{
		cage.GET("/all", r.cageController.GetAllCage)
		cage.GET("/show/:id", r.cageController.ShowCage)
	}

	cageRoutes := route.Group("api/cage", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//cageRoutes.GET("/all", r.cageController.GetAllCage)
		//cageRoutes.GET("/show/:id", r.cageController.ShowCage)
		cageRoutes.POST("/add", r.cageController.CreateCage)
		cageRoutes.PUT("/update", r.cageController.UpdateCage)
		cageRoutes.DELETE("/delete/:id", r.cageController.DeleteCage)
	}

	cageType := route.Group("api/cageType")
	{
		cageType.GET("/all", r.cageTypeController.GetAllCageType)
		cageType.GET("/show/:id", r.cageTypeController.ShowCageType)
	}

	cageTypeRoutes := route.Group("api/cageType", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//cageTypeRoutes.GET("/all", r.cageTypeController.GetAllCageType)
		//cageTypeRoutes.GET("/show/:id", r.cageTypeController.ShowCageType)
		cageTypeRoutes.POST("/add", r.cageTypeController.CreateCageType)
		cageTypeRoutes.PUT("/update", r.cageTypeController.UpdateCageType)
		cageTypeRoutes.DELETE("/delete/:id", r.cageTypeController.DeleteCageType)
	}

	cageCategory := route.Group("api/cageCategory")
	{
		cageCategory.GET("/all", r.cageCategoryController.GetAllCageCategory)
		cageCategory.GET("/show/:id", r.cageCategoryController.ShowCageCategory)
	}

	cageCategoryRoutes := route.Group("api/cageCategory", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		cageCategoryRoutes.POST("/add", r.cageCategoryController.CreateCageCategory)
		cageCategoryRoutes.PUT("/update", r.cageCategoryController.UpdateCageCategory)
		cageCategoryRoutes.DELETE("/delete/:id", r.cageCategoryController.DeleteCageCategory)
	}

	cageDetail := route.Group("api/cageDetail")
	{
		cageDetail.GET("/all", r.cageDetailController.GetAllCageDetail)
		cageDetail.GET("/show/:id", r.cageDetailController.ShowCageDetail)
	}

	cageDetailRoutes := route.Group("api/cageDetail", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//cageDetailRoutes.GET("/all", r.cageDetailController.GetAllCageDetail)
		//cageDetailRoutes.GET("/show/:id", r.cageDetailController.ShowCageDetail)
		cageDetailRoutes.POST("/add", r.cageDetailController.CreateCageDetail)
		cageDetailRoutes.PUT("/update", r.cageDetailController.UpdateCageDetail)
		cageDetailRoutes.PUT("/status", r.cageDetailController.UpdateCageDetailStatus)
		cageDetailRoutes.DELETE("/delete/:id", r.cageDetailController.DeleteCageDetail)
	}

	group := route.Group("api/group")
	{
		group.GET("/all", r.groupController.GetAllGroup)
		group.GET("/show/:id", r.groupController.ShowGroup)
	}

	groupRoutes := route.Group("api/group", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		groupRoutes.POST("/add", r.groupController.CreateGroup)
		groupRoutes.PUT("/update", r.groupController.UpdateGroup)
		groupRoutes.DELETE("/delete/:id", r.groupController.DeleteGroup)
	}

	groupDetail := route.Group("api/groupDetail")
	{
		groupDetail.GET("/all", r.groupDetailController.GetAllGroupDetail)
		groupDetail.GET("/show/:id", r.groupDetailController.ShowGroupDetail)
	}

	groupDetailRoutes := route.Group("api/groupDetail", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//groupDetailRoutes.GET("/all", r.groupDetailController.GetAllGroupDetail)
		//groupDetailRoutes.GET("/show/:id", r.groupDetailController.ShowGroupDetail)
		groupDetailRoutes.POST("/add", r.groupDetailController.CreateGroupDetail)
		groupDetailRoutes.PUT("/update", r.groupDetailController.UpdateGroupDetail)
		groupDetailRoutes.DELETE("/delete/:id", r.groupDetailController.DeleteGroupDetail)
	}

	services := route.Group("api/service")
	{
		services.GET("/all", r.serviceController.GetAllService)
		services.GET("/show/:id", r.serviceController.ShowService)
	}

	serviceRoutes := route.Group("api/service", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//serviceRoutes.GET("/all", r.serviceController.GetAllService)
		//serviceRoutes.GET("/show/:id", r.serviceController.ShowService)
		serviceRoutes.POST("/add", r.serviceController.CreateService)
		serviceRoutes.PUT("/update", r.serviceController.UpdateService)
		serviceRoutes.DELETE("/delete/:id", r.serviceController.DeleteService)
	}

	serviceDetail := route.Group("api/serviceDetail")
	{
		serviceDetail.GET("/all", r.serviceDetailController.GetAllServiceDetail)
		serviceDetail.GET("/show/:id", r.serviceDetailController.ShowServiceDetail)
	}

	serviceDetailRoutes := route.Group("api/serviceDetail", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//serviceDetailRoutes.GET("/all", r.serviceDetailController.GetAllServiceDetail)
		//serviceDetailRoutes.GET("/show/:id", r.serviceDetailController.ShowServiceDetail)
		serviceDetailRoutes.POST("/add", r.serviceDetailController.CreateServiceDetail)
		serviceDetailRoutes.PUT("/update", r.serviceDetailController.UpdateServiceDetail)
		serviceDetailRoutes.DELETE("/delete/:id", r.serviceDetailController.DeleteServiceDetail)
	}

	product := route.Group("api/product")
	{
		product.GET("/all", r.productController.GetAllProduct)
		product.GET("/show/:id", r.productController.ShowProduct)
	}

	productRoutes := route.Group("api/product", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//productRoutes.GET("/all", r.productController.GetAllProduct)
		//productRoutes.GET("/show/:id", r.productController.ShowProduct)
		productRoutes.POST("/add", r.productController.CreateProduct)
		productRoutes.PUT("/update", r.productController.UpdateProduct)
		productRoutes.PUT("/status", r.productController.UpdateProductStatus)
		productRoutes.DELETE("/delete/:id", r.productController.DeleteProduct)
	}

	reservation := route.Group("api/reservation", middleware.CheckRoles(r.jwtService, r.userService, r.redisService, "Admin", "Customer"))
	{
		reservation.GET("/all", r.reservationController.GetAllReservation)
		reservation.GET("/show/:id", r.reservationController.ShowReservation)
	}

	reservationCustomer := route.Group("api/reservationCustomer", middleware.CheckRoles(r.jwtService, r.userService, r.redisService, "Admin", "Customer"))
	{
		reservationCustomer.POST("/add", r.reservationController.CreateReservation)
	}

	reservationRoutes := route.Group("api/reservation", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//reservationRoutes.GET("/all", r.reservationController.GetAllReservation)
		reservationRoutes.PUT("/paymentStatus", r.reservationController.UpdatePaymentStatus)
		reservationRoutes.PUT("/reservationStatus", r.reservationController.UpdateReservationStatus)
		reservationRoutes.PUT("/checkinStatus", r.reservationController.UpdateCheckInStatus)
		//reservationRoutes.GET("/show/:id", r.reservationController.ShowReservation)
		reservationRoutes.PUT("/update", r.reservationController.UpdateReservation)
		reservationRoutes.DELETE("/delete/:id", r.reservationController.DeleteReservation)
	}

	reservationDetail := route.Group("api/reservationDetail", middleware.CheckRoles(r.jwtService, r.userService, r.redisService, "Admin", "Customer"))
	{
		reservationDetail.GET("/all", r.reservationDetailController.GetAllReservationDetail)
		reservationDetail.GET("/show/:id", r.reservationDetailController.ShowReservationDetail)
	}

	reservationDetailRoutes := route.Group("api/reservation_detail", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//reservationDetailRoutes.GET("/all", r.reservationDetailController.GetAllReservationDetail)
		//reservationDetailRoutes.GET("/show/:id", r.reservationDetailController.ShowReservationDetail)
		reservationDetailRoutes.POST("/add", r.reservationDetailController.CreateReservationDetail)
		reservationDetailRoutes.PUT("/update", r.reservationDetailController.UpdateReservationDetail)
		reservationDetailRoutes.DELETE("/delete/:id", r.reservationDetailController.DeleteReservationDetail)
	}

	reservationCondition := route.Group("api/reservationCondition", middleware.CheckRoles(r.jwtService, r.userService, r.redisService, "Admin", "Customer"))
	{
		reservationCondition.GET("/all", r.reservationConditionController.GetAllReservationCondition)
		reservationCondition.GET("/show/:id", r.reservationConditionController.ShowReservationCondition)
	}

	reservationConditionRoutes := route.Group("api/reservation_condition", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//reservationConditionRoutes.GET("/all", r.reservationConditionController.GetAllReservationCondition)
		//reservationConditionRoutes.GET("/show/:id", r.reservationConditionController.ShowReservationCondition)
		reservationConditionRoutes.POST("/add", r.reservationConditionController.CreateReservationCondition)
		reservationConditionRoutes.PUT("/update", r.reservationConditionController.UpdateReservationCondition)
		reservationConditionRoutes.DELETE("/delete/:id", r.reservationConditionController.DeleteReservationCondition)
	}

	reservationInventory := route.Group("api/reservation_inventory", middleware.CheckRoles(r.jwtService, r.userService, r.redisService, "Admin", "Customer"))
	{
		reservationInventory.GET("/all", r.reservationInventoryController.GetAllReservationInventory)
		reservationInventory.GET("/show/:id", r.reservationInventoryController.ShowReservationInventory)
	}

	reservationInventoryRoutes := route.Group("api/reservation_inventory", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//reservationInventoryRoutes.GET("/all", r.reservationInventoryController.GetAllReservationInventory)
		//reservationInventoryRoutes.GET("/show/:id", r.reservationInventoryController.ShowReservationInventory)
		reservationInventoryRoutes.POST("/add", r.reservationInventoryController.CreateReservationInventory)
		reservationInventoryRoutes.PUT("/update", r.reservationInventoryController.UpdateReservationInventory)
		reservationInventoryRoutes.DELETE("/delete:id", r.reservationInventoryController.DeleteReservationInventory)
	}

	reservationService := route.Group("api/reservation_service", middleware.CheckRoles(r.jwtService, r.userService, r.redisService, "Admin", "Customer"))
	{
		reservationService.GET("/all", r.reservationServiceController.GetAllReservationService)
		reservationService.GET("/show/:id", r.reservationServiceController.ShowReservationService)
	}

	reservationServiceRoutes := route.Group("api/reservation_service", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//reservationServiceRoutes.GET("/all", r.reservationServiceController.GetAllReservationService)
		//reservationServiceRoutes.GET("/show/:id", r.reservationServiceController.ShowReservationService)
		reservationServiceRoutes.POST("/add", r.reservationServiceController.CreateReservationService)
		reservationServiceRoutes.PUT("/update", r.reservationServiceController.UpdateReservationService)
		reservationServiceRoutes.DELETE("/delete/:id", r.reservationServiceController.DeleteReservationService)
	}

	reservationProduct := route.Group("api/reservation_product", middleware.CheckRoles(r.jwtService, r.userService, r.redisService, "Admin", "Customer"))
	{
		reservationProduct.GET("/all", r.reservationProductController.GetAllReservationProduct)
		reservationProduct.GET("/show/:id", r.reservationProductController.ShowReservationProduct)
	}

	reservationProductRoutes := route.Group("api/reservation_product", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//reservationProductRoutes.GET("/all", r.reservationProductController.GetAllReservationProduct)
		//reservationProductRoutes.GET("/show/:id", r.reservationProductController.ShowReservationProduct)
		reservationProductRoutes.POST("/add", r.reservationProductController.CreateReservationProduct)
		reservationProductRoutes.PUT("/update", r.reservationProductController.UpdateReservationProduct)
		reservationProductRoutes.DELETE("/delete/:id", r.reservationProductController.DeleteReservationProduct)
	}

	rate := route.Group("api/rate", middleware.CheckRoles(r.jwtService, r.userService, r.redisService, "Admin", "Customer"))
	{
		rate.GET("/all", r.rateController.GetAllRate)
		rate.GET("/show/:id", r.rateController.ShowRate)
	}

	rateRoutes := route.Group("api/rate", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Customer"))
	{
		//rateRoutes.GET("/all", r.rateController.GetAllRate)
		//rateRoutes.GET("/show:id", r.rateController.ShowRate)
		rateRoutes.POST("/add", r.rateController.CreateRate)
		//rateRoutes.PUT("/update", r.rateController.UpdateRate)
		//rateRoutes.DELETE("/delete/:id", r.rateController.DeleteRate)
	}

	response := route.Group("api/response", middleware.CheckRoles(r.jwtService, r.userService, r.redisService, "Admin", "Customer"))
	{
		response.GET("/all", r.responseController.GetAllResponse)
		response.GET("/show/:id", r.responseController.ShowResponse)
	}

	responseRoutes := route.Group("api/response", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		//responseRoutes.GET("/all", r.responseController.GetAllResponse)
		//responseRoutes.GET("/show/:id", r.responseController.ShowResponse)
		responseRoutes.POST("/add", r.responseController.CreateResponse)
		responseRoutes.PUT("/update", r.responseController.UpdateResponse)
		responseRoutes.DELETE("/delete/:id", r.responseController.DeleteResponse)
	}
}

func NewRoute(
	userControllers controller.UserController,
	authControllers controller.AuthController,
	requestController controller.RequestController,
	hotelController controller.HotelController,
	provinceController controller.ProvinceController,
	cityController controller.CityController,
	districtController controller.DistrictController,
	classController controller.ClassController,
	categoryController controller.CategoryController,
	speciesController controller.SpeciesController,
	petController controller.PetController,
	groupController controller.GroupController,
	groupDetailController controller.GroupDetailController,
	cageCategoryController controller.CageCategoryController,
	cageTypeController controller.CageTypeController,
	cageDetailController controller.CageDetailController,
	cageController controller.CageController,
	serviceController controller.ServiceController,
	serviceDetailController controller.ServiceDetailController,
	productController controller.ProductController,
	reservationController controller.ReservationController,
	reservationDetailController controller.ReservationDetailController,
	reservationInventoryController controller.ReservationInventoryController,
	reservationProductController controller.ReservationProductController,
	reservationServiceController controller.ReservationServiceController,
	reservationConditionController controller.ReservationConditionController,
	rateController controller.RateController,
	responseController controller.ResponseController,
	jwtServices service.JWTService,
	redisServices service.RedisService,
	userServices service.UserService,
) Route {
	return &route{
		userController:                 userControllers,
		authController:                 authControllers,
		requestController:              requestController,
		hotelController:                hotelController,
		provinceController:             provinceController,
		cityController:                 cityController,
		districtController:             districtController,
		classController:                classController,
		categoryController:             categoryController,
		speciesController:              speciesController,
		petController:                  petController,
		groupController:                groupController,
		groupDetailController:          groupDetailController,
		cageCategoryController:         cageCategoryController,
		cageTypeController:             cageTypeController,
		cageDetailController:           cageDetailController,
		cageController:                 cageController,
		serviceController:              serviceController,
		serviceDetailController:        serviceDetailController,
		productController:              productController,
		reservationController:          reservationController,
		reservationDetailController:    reservationDetailController,
		reservationProductController:   reservationProductController,
		reservationServiceController:   reservationServiceController,
		reservationInventoryController: reservationInventoryController,
		reservationConditionController: reservationConditionController,
		rateController:                 rateController,
		responseController:             responseController,
		jwtService:                     jwtServices,
		redisService:                   redisServices,
		userService:                    userServices,
	}
}
