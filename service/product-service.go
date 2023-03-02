package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// ProductService is a contract of what productService can do
type ProductService interface {
	GetAllProduct(hotelID string, filterPage dto.FilterPagination) ([]entity.Product, dto.Pagination, error)
	CreateProduct(product dto.CreateProduct) (entity.Product, error)
	UpdateProduct(product dto.UpdateProduct) (entity.Product, error)
	DeleteProduct(productID string, userHotelID string) (error, interface{})
	ShowProduct(productID string, userHotelID string) (entity.Product, error, interface{})
}

type productService struct {
	productRepository repository.ProductRepository
}

// NewProductService creates a new instance of ProductService
func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepo,
	}
}

func (u *productService) CreateProduct(product dto.CreateProduct) (entity.Product, error) {
	productToCreate := entity.Product{}
	errMap := smapping.FillStruct(&productToCreate, smapping.MapFields(&product))
	if errMap != nil {
		return productToCreate, errMap
	}
	updatedProduct, err := u.productRepository.InsertProduct(productToCreate)
	return updatedProduct, err
}
func (u *productService) UpdateProduct(product dto.UpdateProduct) (entity.Product, error) {
	productToUpdate := entity.Product{}
	errMap := smapping.FillStruct(&productToUpdate, smapping.MapFields(&product))
	if errMap != nil {
		return productToUpdate, errMap
	}
	updatedProduct, err := u.productRepository.UpdateProduct(productToUpdate)
	return updatedProduct, err
}
func (u *productService) DeleteProduct(productID string, userHotelID string) (error, interface{}) {
	product, err := u.productRepository.FindProductByID(productID)
	if product.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	errDel := u.productRepository.DeleteProduct(product)
	return errDel, nil
}
func (u *productService) ShowProduct(productID string, userHotelID string) (entity.Product, error, interface{}) {
	result, err := u.productRepository.FindProductByID(productID)
	if result.HotelID != userHotelID {
		return entity.Product{}, nil, false
	}
	return result, err, nil
}
func (u *productService) GetAllProduct(hotelID string, filterPage dto.FilterPagination) ([]entity.Product, dto.Pagination, error) {
	admins, pagination, err := u.productRepository.GetAllProduct(hotelID, filterPage)
	return admins, pagination, err
}
