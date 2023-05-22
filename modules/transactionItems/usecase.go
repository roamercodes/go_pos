package transactions

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetAllTransactionItems() ([]TransactionItem, error) {
	transactionItems, err := usecase.Repo.GetAllTransactionItems()
	return transactionItems, err
}

func (usecase Usecase) GetTransactionItemById(id string) (*TransactionItem, error) {
	transactionItems, err := usecase.Repo.GetTransactionById(id)
	return transactionItems, err
}

func (usecase Usecase) CreateTransactionItem(transaction TransactionItem) error {
	err := usecase.Repo.CreateTransactionItem(transaction)
	return err
}

func (usecase Usecase) UpdateTransactionItemById(id string, transactionItem TransactionItem) error {
	err := usecase.Repo.UpdateTransactionItemById(id, transactionItem)
	return err
}

func (usecase Usecase) DeleteTransactionItemById(id string) error {
	err := usecase.Repo.DeleteTransactionItemById(id)
	return err
}