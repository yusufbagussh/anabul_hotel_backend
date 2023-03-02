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

	requestRoutes := route.Group("api/request", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		requestRoutes.GET("/all", r.requestController.GetAllRequest)
		requestRoutes.GET("/:id", r.requestController.ShowRequest)
		requestRoutes.POST("/add", r.requestController.CreateRequest)
		requestRoutes.PUT("/update", r.requestController.UpdateRequest)
		requestRoutes.DELETE("/:id", r.requestController.DeleteRequest)
	}

	userRoutes := route.Group("api/user", middleware.AuthorizeJWT(r.jwtService, r.userService, r.redisService))
	{
		userRoutes.GET("/profile", r.userController.GetProfile)
		//userRoutes.GET("/notif", r.userController.SendNotif)
		userRoutes.PUT("/profile", r.userController.UpdateProfile)
		userRoutes.PUT("/changepass", r.userController.ChangePassword)
	}

	adminRoutes := route.Group("api/admin", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		adminRoutes.GET("/all", r.userController.GetAdmin)
		adminRoutes.GET("/:id", r.userController.ShowAdmin)
		adminRoutes.POST("/add", r.userController.CreateUser)
		adminRoutes.PUT("/update", r.userController.UpdateAdmin)
		adminRoutes.DELETE("/:id", r.userController.DeleteAdmin)
	}

	hotelRoutes := route.Group("api/hotel", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		hotelRoutes.GET("/all", r.hotelController.GetAllHotel)
		hotelRoutes.GET("/:id", r.hotelController.ShowHotel)
		hotelRoutes.POST("/add", r.hotelController.CreateHotel)
		hotelRoutes.POST("/admin", r.hotelController.CreateHotelAdmin)
		hotelRoutes.PUT("/update", r.hotelController.UpdateHotel)
		hotelRoutes.DELETE("/:id", r.hotelController.DeleteHotel)
	}

	hotelAdminRoutes := route.Group("api/profilehotel", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		hotelAdminRoutes.GET("/profile", r.hotelController.GetProfileHotel)
		hotelAdminRoutes.PUT("/profile", r.hotelController.UpdateProfileHotel)
	}

	classRoutes := route.Group("api/class")
	{
		classRoutes.GET("/all", r.classController.GetAllClass)
		classRoutes.GET("/:id", r.classController.ShowClass)
		classRoutes.POST("/add", r.classController.CreateClass)
		classRoutes.PUT("/update", r.classController.UpdateClass)
		classRoutes.DELETE("/:id", r.classController.DeleteClass)
	}

	class := route.Group("api/class", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		class.POST("/add", r.categoryController.CreateCategory)
		class.PUT("/update", r.categoryController.UpdateCategory)
		class.DELETE("/:id", r.categoryController.DeleteCategory)
	}

	categoryRoutes := route.Group("api/category", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Super Admin"))
	{
		categoryRoutes.GET("/all", r.categoryController.GetAllCategory)
		categoryRoutes.GET("/:id", r.categoryController.ShowCategory)
		categoryRoutes.POST("/add", r.categoryController.CreateCategory)
		categoryRoutes.PUT("/update", r.categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", r.categoryController.DeleteCategory)
	}

	speciesRoutes := route.Group("api/species", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		speciesRoutes.GET("/all", r.speciesController.GetSpecies)
		speciesRoutes.GET("/:id", r.speciesController.ShowSpecies)
		speciesRoutes.POST("/add", r.speciesController.CreateSpecies)
		speciesRoutes.PUT("/update", r.speciesController.UpdateSpecies)
		speciesRoutes.DELETE("/:id", r.speciesController.DeleteSpecies)
	}

	petRoutes := route.Group("api/pet", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Customer"))
	{
		petRoutes.GET("/", r.petController.GetPet)
		petRoutes.GET("/:id", r.petController.ShowPet)
		petRoutes.POST("/", r.petController.CreatePet)
		petRoutes.PUT("/", r.petController.UpdatePet)
		petRoutes.DELETE("/:id", r.petController.DeletePet)
	}

	staffRoutes := route.Group("api/staff", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		staffRoutes.GET("/", r.userController.GetStaff)
		staffRoutes.GET("/:id", r.userController.ShowStaff)
		staffRoutes.POST("/", r.userController.CreateUser)
		staffRoutes.PUT("/", r.userController.UpdateStaff)
		staffRoutes.DELETE("/:id", r.userController.DeleteStaff)
	}

	cageRoutes := route.Group("api/cage", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		cageRoutes.GET("/", r.cageController.GetAllCage)
		cageRoutes.GET("/:id", r.cageController.ShowCage)
		cageRoutes.POST("/", r.cageController.CreateCage)
		cageRoutes.PUT("/", r.cageController.UpdateCage)
		cageRoutes.DELETE("/:id", r.cageController.DeleteCage)
	}

	cageTypeRoutes := route.Group("api/cageType", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		cageTypeRoutes.GET("/", r.cageTypeController.GetAllCageType)
		cageTypeRoutes.GET("/:id", r.cageTypeController.ShowCageType)
		cageTypeRoutes.POST("/", r.cageTypeController.CreateCageType)
		cageTypeRoutes.PUT("/", r.cageTypeController.UpdateCageType)
		cageTypeRoutes.DELETE("/:id", r.cageTypeController.DeleteCageType)
	}

	cageCategoryRoutes := route.Group("api/cageCategory", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		cageCategoryRoutes.GET("/", r.cageCategoryController.GetAllCageCategory)
		cageCategoryRoutes.GET("/:id", r.cageCategoryController.ShowCageCategory)
		cageCategoryRoutes.POST("/", r.cageCategoryController.CreateCageCategory)
		cageCategoryRoutes.PUT("/", r.cageCategoryController.UpdateCageCategory)
		cageCategoryRoutes.DELETE("/:id", r.cageCategoryController.DeleteCageCategory)
	}

	cageDetailRoutes := route.Group("api/cageDetail", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		cageDetailRoutes.GET("/", r.cageDetailController.GetAllCageDetail)
		cageDetailRoutes.GET("/:id", r.cageDetailController.ShowCageDetail)
		cageDetailRoutes.POST("/", r.cageDetailController.CreateCageDetail)
		cageDetailRoutes.PUT("/", r.cageDetailController.UpdateCageDetail)
		cageDetailRoutes.DELETE("/:id", r.cageDetailController.DeleteCageDetail)
	}

	groupRoutes := route.Group("api/group", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		groupRoutes.GET("/", r.groupController.GetAllGroup)
		groupRoutes.GET("/:id", r.groupController.ShowGroup)
		groupRoutes.POST("/", r.groupController.CreateGroup)
		groupRoutes.PUT("/", r.groupController.UpdateGroup)
		groupRoutes.DELETE("/:id", r.groupController.DeleteGroup)
	}

	groupDetailRoutes := route.Group("api/groupDetail", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		groupDetailRoutes.GET("/", r.groupDetailController.GetAllGroupDetail)
		groupDetailRoutes.GET("/:id", r.groupDetailController.ShowGroupDetail)
		groupDetailRoutes.POST("/", r.groupDetailController.CreateGroupDetail)
		groupDetailRoutes.PUT("/", r.groupDetailController.UpdateGroupDetail)
		groupDetailRoutes.DELETE("/:id", r.groupDetailController.DeleteGroupDetail)
	}

	serviceRoutes := route.Group("api/service", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		serviceRoutes.GET("/", r.serviceController.GetAllService)
		serviceRoutes.GET("/:id", r.serviceController.ShowService)
		serviceRoutes.POST("/", r.serviceController.CreateService)
		serviceRoutes.PUT("/", r.serviceController.UpdateService)
		serviceRoutes.DELETE("/:id", r.serviceController.DeleteService)
	}

	serviceDetailRoutes := route.Group("api/serviceDetail", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		serviceDetailRoutes.GET("/", r.serviceDetailController.GetAllServiceDetail)
		serviceDetailRoutes.GET("/:id", r.serviceDetailController.ShowServiceDetail)
		serviceDetailRoutes.POST("/", r.serviceDetailController.CreateServiceDetail)
		serviceDetailRoutes.PUT("/", r.serviceDetailController.UpdateServiceDetail)
		serviceDetailRoutes.DELETE("/:id", r.serviceDetailController.DeleteServiceDetail)
	}

	productRoutes := route.Group("api/product", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	{
		productRoutes.GET("/", r.productController.GetAllProduct)
		productRoutes.GET("/:id", r.productController.ShowProduct)
		productRoutes.POST("/", r.productController.CreateProduct)
		productRoutes.PUT("/", r.productController.UpdateProduct)
		productRoutes.DELETE("/:id", r.productController.DeleteProduct)
	}

	reservationRoutes := route.Group("api/reservation", middleware.AuthorizeJWT(r.jwtService, r.userService, r.redisService))
	{
		reservationRoutes.GET("/", r.reservationController.GetAllReservation)
		reservationRoutes.GET("/:id", r.reservationController.ShowReservation)
		reservationRoutes.POST("/", r.reservationController.CreateReservation)
		reservationRoutes.PUT("/", r.reservationController.UpdateReservation)
		reservationRoutes.DELETE("/:id", r.reservationController.DeleteReservation)
	}

	//reservationDetailRoutes := route.Group("api/reservation_detail", middleware.AuthorizeJWT(r.jwtService, r.userService, r.redisService))
	//{
	//	reservationDetailRoutes.GET("/", r.reservationDetailController.GetAllReservationDetail)
	//	reservationDetailRoutes.GET("/:id", r.reservationDetailController.ShowReservationDetail)
	//	reservationDetailRoutes.POST("/", r.reservationDetailController.CreateReservationDetail)
	//	reservationDetailRoutes.PUT("/", r.reservationDetailController.UpdateReservationDetail)
	//	reservationDetailRoutes.DELETE("/:id", r.reservationDetailController.DeleteReservationDetail)
	//}

	reservationInventoryRoutes := route.Group("api/reservation_inventory", middleware.AuthorizeJWT(r.jwtService, r.userService, r.redisService))
	{
		reservationInventoryRoutes.GET("/", r.reservationInventoryController.GetAllReservationInventory)
		reservationInventoryRoutes.GET("/:id", r.reservationInventoryController.ShowReservationInventory)
		reservationInventoryRoutes.POST("/", r.reservationInventoryController.CreateReservationInventory)
		reservationInventoryRoutes.PUT("/", r.reservationInventoryController.UpdateReservationInventory)
		reservationInventoryRoutes.DELETE("/:id", r.reservationInventoryController.DeleteReservationInventory)
	}

	//reservationServiceRoutes := route.Group("api/reservation_service", middleware.AuthorizeJWT(r.jwtService, r.userService, r.redisService))
	//{
	//	reservationServiceRoutes.GET("/", r.reservationServiceController.GetAllReservationService)
	//	reservationServiceRoutes.GET("/:id", r.reservationServiceController.ShowReservationService)
	//	reservationServiceRoutes.POST("/", r.reservationServiceController.CreateReservationService)
	//	reservationServiceRoutes.PUT("/", r.reservationServiceController.UpdateReservationService)
	//	reservationServiceRoutes.DELETE("/:id", r.reservationServiceController.DeleteReservationService)
	//}
	//
	//reservationProductRoutes := route.Group("api/reservation_product", middleware.AuthorizeJWT(r.jwtService, r.userService, r.redisService))
	//{
	//	reservationProductRoutes.GET("/", r.reservationProductController.GetAllReservationProduct)
	//	reservationProductRoutes.GET("/:id", r.reservationProductController.ShowReservationProduct)
	//	reservationProductRoutes.POST("/", r.reservationProductController.CreateReservationProduct)
	//	reservationProductRoutes.PUT("/", r.reservationProductController.UpdateReservationProduct)
	//	reservationProductRoutes.DELETE("/:id", r.reservationProductController.DeleteReservationProduct)
	//}

	//rateRoutes := route.Group("api/rate", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Customer"))
	//{
	//	rateRoutes.GET("/", r.rateController.GetAllRate)
	//	rateRoutes.GET("/:id", r.rateController.ShowRate)
	//	rateRoutes.POST("/", r.rateController.CreateRate)
	//	rateRoutes.PUT("/", r.rateController.UpdateRate)
	//	rateRoutes.DELETE("/:id", r.rateController.DeleteRate)
	//}
	//
	//responseRoutes := route.Group("api/response", middleware.CheckRole(r.jwtService, r.userService, r.redisService, "Admin"))
	//{
	//	responseRoutes.GET("/", r.responseController.GetAllResponse)
	//	responseRoutes.GET("/:id", r.responseController.ShowResponse)
	//	responseRoutes.POST("/", r.responseController.CreateResponse)
	//	responseRoutes.PUT("/", r.responseController.UpdateResponse)
	//	responseRoutes.DELETE("/:id", r.responseController.DeleteResponse)
	//}
}

func NewRoute(
	userControllers controller.UserController,
	authControllers controller.AuthController,
	requestController controller.RequestController,
	hotelController controller.HotelController,
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
	//reservationDetailController controller.ReservationDetailController,
	reservationInventoryController controller.ReservationInventoryController,
	//reservationProductController controller.ReservationProductController,
	//reservationServiceController controller.ReservationServiceController,
	//rateController controller.RateController,
	//responseController controller.ResponseController,
	jwtServices service.JWTService,
	redisServices service.RedisService,
	userServices service.UserService,
) Route {
	return &route{
		userController:          userControllers,
		authController:          authControllers,
		requestController:       requestController,
		hotelController:         hotelController,
		classController:         classController,
		categoryController:      categoryController,
		speciesController:       speciesController,
		petController:           petController,
		groupController:         groupController,
		groupDetailController:   groupDetailController,
		cageCategoryController:  cageCategoryController,
		cageTypeController:      cageTypeController,
		cageDetailController:    cageDetailController,
		cageController:          cageController,
		serviceController:       serviceController,
		serviceDetailController: serviceDetailController,
		productController:       productController,
		reservationController:   reservationController,
		//reservationDetailController:    reservationDetailController,
		//reservationProductController:   reservationProductController,
		//reservationServiceController:   reservationServiceController,
		reservationInventoryController: reservationInventoryController,
		//rateController:                 rateController,
		//responseController:             responseController,
		jwtService:   jwtServices,
		redisService: redisServices,
		userService:  userServices,
	}
}
