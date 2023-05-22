package transactions

import (
	"time"
	"project-week1/modules/products"
	// "project-week1/modules/login"
)

type Transaction struct {
	Id			int					`gorm:"primaryKey; column:id" json:"id"`
	Timestamp	time.Time			`json:"timestamp"`
	Total		int					`gorm:"total" json:"total"`
	AdminId		int					`gorm:"Foreignkey:admin_id;association_foreignkey:id;" json:"admin_id"`
	Admin		User			`json:"admin"`
	Items		[]TransactionItem	`json:"items"`
}

type User struct {
	Id		int
	Fullname	string		
}

type TransactionItem struct {
	Id				int 	`gorm:"primaryKey; column:id" json:"id"`
	TransactionId	int 	`gorm:"Foreignkey:transaction_id;association_foreignkey:id;" json:"transaction_id"`
	ProductId		int 	`gorm:"Foreignkey:product_id;association_foreignkey:id;" json:"product_id"`
	Quantity		int		`gorm:"int(11)" json:"quantity"`
	Price			int		`gorm:"int(11)" json:"price"`
	Product			*products.Product `json:"product"`
}

type TransactionItemRequest struct {
	ProductId int `json:"product_id"`
	Quantity int `json:"quantity"`
}

type CreateTransactionRequest struct {
	Items []TransactionItem 
}