package products

import (
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllProducts() ([]Product, error) {
	var products []Product
	result := repo.DB.Find(&products)
	return products, result.Error
}

func (repo Repository) GetProductByIdf(id string) (*Product, error) {
	var product Product
	result := repo.DB.First(&product, id)
	return &product, result.Error
}

func (repo Repository) CreateProduct(product Product) error {
	result := repo.DB.Create(&product)
	return result.Error
}

func (repo Repository) UpdateProductById(id int, product *Product) error {
	result := repo.DB.Model(&Product{}).Where("id = ?", id).Updates(product)
	return result.Error
}

//
func (repo Repository) EditProduct(id int, product *Product) error{
	result := repo.DB.Where(id).Updates(&product)

	return result.Error
}

//
func (repo Repository) InactiveProduct(id int) error{
	result := repo.DB.Model(&Product{}).Where("id = ?", id).Updates(Product{Deleted_at: time.Now()})
	return result.Error
}

func (repo Repository) ActiveProduct(id int) error{
	result := repo.DB.Model(&Product{}).Where("id = ?", id).Updates(map[string]interface{}{"Deleted_at": nil})
	return result.Error
}

//
func (repo Repository) GetProductById(id int) (*Product, error){
	var product *Product
	result := repo.DB.First(&product, id)

	return product, result.Error
}
