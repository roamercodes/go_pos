package transactions

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllTransactionItems() ([]TransactionItem, error) {
	var transactionItems []TransactionItem
	result := repo.DB.Find(&transactionItems)
	return transactionItems, result.Error
}

func (repo Repository) GetTransactionById(id string) (*TransactionItem, error) {
	var transactionItem *TransactionItem
	result := repo.DB.First(&transactionItem, id)
	return transactionItem, result.Error
}

func (repo Repository) CreateTransactionItem(transactionItem TransactionItem) error {
	result := repo.DB.Create(&transactionItem)
	return result.Error
}

func (repo Repository) UpdateTransactionItemById(id string, transactionItem TransactionItem) error {
	result := repo.DB.Model(&TransactionItem{}).Where("id = ?", id).Updates(transactionItem)
	return result.Error
}

func (repo Repository) DeleteTransactionItemById(id string) error {
	result := repo.DB.Delete(&TransactionItem{}, id)
	return result.Error
}