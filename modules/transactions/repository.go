package transactions

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllTransactions() ([]Transaction, error) {
	var transactions []Transaction
	// result := repo.DB.Find(&transactions)
	result := repo.DB.Preload("Items").Find(&transactions)
	return transactions, result.Error
}

func (repo Repository) GetTransactionById(id string) (*Transaction, error) {
	var transaction *Transaction
	result := repo.DB.Preload("Items").First(&transaction, id)
	return transaction, result.Error
}

// n
func (repo Repository) GetById(id int) (*Transaction, error){
	var transaction *Transaction
	result := repo.DB.Preload("Admin").Preload("Items.Product").First(&transaction, id)

	return transaction, result.Error
}

// n
func (repo Repository) Create(transaction *Transaction) error{
	result := repo.DB.Create(transaction)

	return result.Error
} 	

// func (repo Repository) CreateTransaction(transaction Transaction) error {
// 	// result := repo.DB.Create(&transaction)
// 	res := repo.DB.Create(&transaction)
// 	return res.Error
// }

func (repo Repository) UpdateTransactionById(id string, transaction Transaction) error {
	result := repo.DB.Model(&Transaction{}).Where("id = ?", id).Updates(transaction)
	return result.Error
}

func (repo Repository) DeleteTransactionById(id string) error {
	result := repo.DB.Delete(&Transaction{}, id)
	return result.Error
}