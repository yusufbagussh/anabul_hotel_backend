package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// ProductRepository is contract what productRepository can do to db
type ProductRepository interface {
	GetAllProduct(hotelID string, filterPagination dto.FilterPagination) ([]entity.Product, dto.Pagination, error)
	InsertProduct(product entity.Product) (entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	FindProductByID(productID string) (entity.Product, error)
	DeleteProduct(product entity.Product) error
	UpdateProductStatus(productStatus dto.UpdateProductStatus) (entity.Product, error)
}

// productConnection adalah func untuk melakukan query data ke tabel product
type productConnection struct {
	connection *gorm.DB
}

func (db *productConnection) DeleteProduct(product entity.Product) error {
	err := db.connection.Where("id_product = ?", product.IDProduct).Delete(&product).Error
	return err
}

func (db *productConnection) UpdateProductStatus(productStatus dto.UpdateProductStatus) (entity.Product, error) {
	var product entity.Product
	err := db.connection.Model(&product).Where("id_reservation = ?", productStatus.IDProduct).Updates(&entity.Product{Status: productStatus.Status}).Error
	db.connection.Find(&product)
	return product, err
}

func (db *productConnection) GetAllProduct(hotelID string, filterPagination dto.FilterPagination) ([]entity.Product, dto.Pagination, error) {
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
	var total int64

	var products []entity.Product
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(products.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
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

	err := query.Where("hotel_id = ?", hotelID).Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&products).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return products, pagination, err
}

// InsertProduct is to add product in database
func (db *productConnection) InsertProduct(product entity.Product) (entity.Product, error) {
	err := db.connection.Save(&product).Error
	db.connection.Find(&product)
	return product, err
}

// UpdateProduct is func to edit product in database
func (db *productConnection) UpdateProduct(product entity.Product) (entity.Product, error) {
	err := db.connection.Where("id_product = ?", product.IDProduct).Updates(&product).Error
	db.connection.Where("id_product = ?", product.IDProduct).Find(&product)
	return product, err
}

// FindProductByID is func to get product by email
func (db *productConnection) FindProductByID(productID string) (entity.Product, error) {
	var product entity.Product
	err := db.connection.Where("id_product = ?", productID).Take(&product).Error
	return product, err
}

// NewProductRepository is creates a new instance of ProductRepository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}
