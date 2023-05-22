package products

import (
	"time"
)

type Product struct {
	Id		int 	`json:"id"`
	Name	string	`json:"name"`
	Price	int		`json:"price"`
	Stock	int		`json:"stock"`
	Deleted_at time.Time `json:"deleted_at"`
}

type StatusRequest struct {
	Status string `json:"status"`
}

type ProductsResponse struct{
	Message string
	Data []Product
}

type ProductResponse struct{
	Message string
	Data *Product
}

type ResponseAddAndEditData struct{
	Message string
	Data Product
}

