package products

import (
	"fmt"
	// "time"
	"strings"
	// "errors"
)

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetAllProducts() ([]Product, error) {
	products, err := usecase.Repo.GetAllProducts()
	return products, err
}

func (usecase Usecase) GetProductByIdf(id string) (*Product, error) {
	products, err := usecase.Repo.GetProductByIdf(id)
	return products, err
}

func (usecase Usecase) CreateProduct(product Product) error {
	err := usecase.Repo.CreateProduct(product)
	return err
}

func (usecase Usecase) UpdateProductById(id int, product *Product) error {
	err := usecase.Repo.UpdateProductById(id, product)
	return err
}

func (usecase Usecase) SoftDeleteInactiveProduct(id int, statusRequest StatusRequest) error {

	product, err := usecase.Repo.GetProductById(id)
	strInactive := "inactive" 
	// strActive := "active"

	if product.Deleted_at.IsZero() == true {
		if strings.EqualFold(statusRequest.Status, strInactive)  {
			err := usecase.Repo.InactiveProduct(id)
			return err
		}
	} else {
		return fmt.Errorf("product already deleted")
	}
	// else if product.Deleted_at.IsZero() == false {
	// 	if strings.EqualFold(statusRequest.Status, strActive)  {
	// 		err := usecase.Repo.ActiveProduct(id)
	// 		return err
	// 	} else {
	// 		fmt.Println("tanggal tidak null")
	// 		return fmt.Errorf("cannot changing the data")
	// 	}
	// }
	return err
}

func (usecase Usecase) SoftDeleteActiveProduct(id int, statusRequest StatusRequest) error {

	product, err := usecase.Repo.GetProductById(id)
	// strInactive := "inactive" 
	strActive := "active"

	// if product.Deleted_at.IsZero() == true {
	// 	if strings.EqualFold(statusRequest.Status, strInactive)  {
	// 		err := usecase.Repo.InactiveProduct(id)
	// 		fmt.Println("tanggal null")
	// 		return err
	// 	} else {
	// 		fmt.Println("tanggal tidak null")
	// 		return fmt.Errorf("cannot changing the data")
	// 	}
	// } 
	if product.Deleted_at.IsZero() == false {
		if strings.EqualFold(statusRequest.Status, strActive)  {
			err := usecase.Repo.ActiveProduct(id)
			return err
		}
	} else {
		return fmt.Errorf("product is not deleted")
	}
	return err
}

// func (usecase Usecase) DeleteProductById(id string) error {
// 	err := usecase.Repo.DeleteProductById(id)
// 	return err
// }

// n
func (usecase Usecase) GetProductById(id int) (*Product, error) {
	product, err := usecase.Repo.GetProductById(id)

	// producttime := product.Deleted_at.IsZero()
	// if producttime == true {
	// 	status = "active"
	// }
	
	return product, err
}
