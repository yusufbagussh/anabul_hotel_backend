package main

import (
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/yusufbagussh/pet_hotel_backend/config"
	"github.com/yusufbagussh/pet_hotel_backend/controller"
	"github.com/yusufbagussh/pet_hotel_backend/database/migration"
	"github.com/yusufbagussh/pet_hotel_backend/database/seeder"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/middleware"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
	"github.com/yusufbagussh/pet_hotel_backend/route"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"gorm.io/gorm"
	"os"
	//"time"
	//"github.com/gin-contrib/cors"
)

var (
	db                             *gorm.DB                                  = config.SetupDatabaseConnection()
	cache                          *redis.Client                             = config.SetupRedisConnection()
	Migrate                        migration.Migrator                        = migration.NewMigration(db)
	Seed                           seeder.Seeder                             = seeder.NewSeeder(db)
	requestRepository              repository.RequestRepository              = repository.NewRequestRepository(db)
	hotelRepository                repository.HotelRepository                = repository.NewHotelRepository(db)
	userRepository                 repository.UserRepository                 = repository.NewUserRepository(db)
	provinceRepository             repository.ProvinceRepository             = repository.NewProvinceRepository(db)
	cityRepository                 repository.CityRepository                 = repository.NewCityRepository(db)
	districtRepository             repository.DistrictRepository             = repository.NewDistrictRepository(db)
	classRepository                repository.ClassRepository                = repository.NewClassRepository(db)
	categoryRepository             repository.CategoryRepository             = repository.NewCategoryRepository(db)
	speciesRepository              repository.SpeciesRepository              = repository.NewSpeciesRepository(db)
	petRepository                  repository.PetRepository                  = repository.NewPetRepository(db)
	groupRepository                repository.GroupRepository                = repository.NewGroupRepository(db)
	groupDetailRepository          repository.GroupDetailRepository          = repository.NewGroupDetailRepository(db)
	cageCategoryRepository         repository.CageCategoryRepository         = repository.NewCageCategoryRepository(db)
	cageTypeRepository             repository.CageTypeRepository             = repository.NewCageTypeRepository(db)
	cageDetailRepository           repository.CageDetailRepository           = repository.NewCageDetailRepository(db)
	cageRepository                 repository.CageRepository                 = repository.NewCageRepository(db)
	serviceRepository              repository.ServiceRepository              = repository.NewServiceRepository(db)
	serviceDetailRepository        repository.ServiceDetailRepository        = repository.NewServiceDetailRepository(db)
	productRepository              repository.ProductRepository              = repository.NewProductRepository(db)
	reservationRepository          repository.ReservationRepository          = repository.NewReservationRepository(db)
	reservationDetailRepository    repository.ReservationDetailRepository    = repository.NewReservationDetailRepository(db)
	reservationProductRepository   repository.ReservationProductRepository   = repository.NewReservationProductRepository(db)
	reservationServiceRepository   repository.ReservationServiceRepository   = repository.NewReservationServiceRepository(db)
	reservationInventoryRepository repository.ReservationInventoryRepository = repository.NewReservationInventoryRepository(db)
	reservationConditionRepository repository.ReservationConditionRepository = repository.NewReservationConditionRepository(db)
	rateRepository                 repository.RateRepository                 = repository.NewRateRepository(db)
	responseRepository             repository.ResponseRepository             = repository.NewResponseRepository(db)
	redisService                   service.RedisService                      = service.NewRedisService(cache)
	jwtService                     service.JWTService                        = service.NewJWTService()
	checkHelper                    helper.CheckHelper                        = helper.NewCheckHelper(userRepository)
	authService                    service.AuthService                       = service.NewAuthService(userRepository, redisService)
	requestService                 service.RequestService                    = service.NewRequestService(requestRepository, userRepository, hotelRepository, checkHelper)
	hotelService                   service.HotelService                      = service.NewHotelService(hotelRepository, checkHelper)
	userService                    service.UserService                       = service.NewUserService(userRepository)
	provinceService                service.ProvinceService                   = service.NewProvinceService(provinceRepository)
	cityService                    service.CityService                       = service.NewCityService(cityRepository)
	districtService                service.DistrictService                   = service.NewDistrictService(districtRepository)
	classService                   service.ClassService                      = service.NewClassService(classRepository)
	categoryService                service.CategoryService                   = service.NewCategoryService(categoryRepository)
	speciesService                 service.SpeciesService                    = service.NewSpeciesService(speciesRepository)
	petService                     service.PetService                        = service.NewPetService(petRepository)
	groupService                   service.GroupService                      = service.NewGroupService(groupRepository)
	groupDetailService             service.GroupDetailService                = service.NewGroupDetailService(groupDetailRepository)
	cageCategoryService            service.CageCategoryService               = service.NewCageCategoryService(cageCategoryRepository)
	cageTypeService                service.CageTypeService                   = service.NewCageTypeService(cageTypeRepository)
	cageDetailService              service.CageDetailService                 = service.NewCageDetailService(cageDetailRepository)
	cageService                    service.CageService                       = service.NewCageService(cageRepository)
	serviceService                 service.ServiceService                    = service.NewServiceService(serviceRepository)
	serviceDetailService           service.ServiceDetailService              = service.NewServiceDetailService(serviceDetailRepository)
	productService                 service.ProductService                    = service.NewProductService(productRepository)
	reservationService             service.ReservationService                = service.NewReservationService(reservationRepository)
	reservationDetailService       service.ReservationDetailService          = service.NewReservationDetailService(reservationDetailRepository)
	reservationServiceService      service.ReservationServiceService         = service.NewReservationServiceService(reservationServiceRepository)
	reservationProductService      service.ReservationProductService         = service.NewReservationProductService(reservationProductRepository)
	reservationInventoryService    service.ReservationInventoryService       = service.NewReservationInventoryService(reservationInventoryRepository)
	reservationConditionService    service.ReservationConditionService       = service.NewReservationConditionService(reservationConditionRepository)
	rateService                    service.RateService                       = service.NewRateService(rateRepository)
	responseService                service.ResponseService                   = service.NewResponseService(responseRepository)
	authController                 controller.AuthController                 = controller.NewAuthController(authService, jwtService, redisService)
	requestController              controller.RequestController              = controller.NewRequestController(requestService, jwtService)
	hotelController                controller.HotelController                = controller.NewHotelController(hotelService, jwtService)
	userController                 controller.UserController                 = controller.NewUserController(userService, jwtService)
	classController                controller.ClassController                = controller.NewClassController(classService, jwtService)
	provinceController             controller.ProvinceController             = controller.NewProvinceController(provinceService, jwtService)
	cityController                 controller.CityController                 = controller.NewCityController(cityService, jwtService)
	districtController             controller.DistrictController             = controller.NewDistrictController(districtService, jwtService)
	categoryController             controller.CategoryController             = controller.NewCategoryController(categoryService, jwtService)
	speciesController              controller.SpeciesController              = controller.NewSpeciesController(speciesService, jwtService)
	petController                  controller.PetController                  = controller.NewPetController(petService, jwtService)
	groupController                controller.GroupController                = controller.NewGroupController(groupService, jwtService)
	groupDetailController          controller.GroupDetailController          = controller.NewGroupDetailController(groupDetailService, jwtService)
	cageCategoryController         controller.CageCategoryController         = controller.NewCageCategoryController(cageCategoryService, jwtService)
	cageTypeController             controller.CageTypeController             = controller.NewCageTypeController(cageTypeService, jwtService)
	cageDetailController           controller.CageDetailController           = controller.NewCageDetailController(cageDetailService, jwtService)
	cageController                 controller.CageController                 = controller.NewCageController(cageService, jwtService)
	serviceController              controller.ServiceController              = controller.NewServiceController(serviceService, jwtService)
	serviceDetailController        controller.ServiceDetailController        = controller.NewServiceDetailController(serviceDetailService, jwtService)
	productController              controller.ProductController              = controller.NewProductController(productService, jwtService)
	reservationController          controller.ReservationController          = controller.NewReservationController(reservationService, jwtService)
	reservationDetailController    controller.ReservationDetailController    = controller.NewReservationDetailController(reservationDetailService, jwtService)
	reservationServiceController   controller.ReservationServiceController   = controller.NewReservationServiceController(reservationServiceService, jwtService)
	reservationProductController   controller.ReservationProductController   = controller.NewReservationProductController(reservationProductService, jwtService)
	reservationInventoryController controller.ReservationInventoryController = controller.NewReservationInventoryController(reservationInventoryService, jwtService)
	reservationConditionController controller.ReservationConditionController = controller.NewReservationConditionController(reservationConditionService, jwtService)
	rateController                 controller.RateController                 = controller.NewRateController(rateService, jwtService)
	responseController             controller.ResponseController             = controller.NewResponseController(responseService, jwtService)
	Routes                         route.Route                               = route.NewRoute(
		userController,
		authController,
		requestController,
		hotelController,
		provinceController,
		cityController,
		districtController,
		classController,
		categoryController,
		speciesController,
		petController,
		groupController,
		groupDetailController,
		cageCategoryController,
		cageTypeController,
		cageDetailController,
		cageController,
		serviceController,
		serviceDetailController,
		productController,
		reservationController,
		reservationDetailController,
		reservationInventoryController,
		reservationProductController,
		reservationServiceController,
		reservationConditionController,
		rateController,
		responseController,
		jwtService,
		redisService,
		userService,
	)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	defer cache.Close()
	if os.Getenv("APP_ENV") == "DEVELOPMENT" {
		Migrate.DropTable()
		Migrate.Migration()
		Seed.Seeder()
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	Routes.Routes(router)
	router.Run()
}
