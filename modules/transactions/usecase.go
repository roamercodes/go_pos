package transactions

import (
	"project-week1/modules/products"
	// "strconv"
	"fmt"
	"time"
	// "net/http"
	// "errors"
	"context"
)

type Usecase struct {
	TransactionRepo Repository
	ProductRepo products.Repository
}

func (usecase Usecase) GetAllTransactions() ([]Transaction, error) {
	transactions, err := usecase.TransactionRepo.GetAllTransactions()
	return transactions, err
}

func (usecase Usecase) GetTransactionById(id int) (*Transaction, error) {
	// transactions, err := usecase.TransactionRepo.GetTransactionById(id)
	transaction, err := usecase.TransactionRepo.GetById(id)
	return transaction, err
}

func (usecase Usecase) CreateTransaction(ctx context.Context, transactionRequest *CreateTransactionRequest) (*Transaction, error)  {
	
	// fmt.Println(ctx.Value("id"))
	// ctx.Value("name")
	id := ctx.Value("id")
	name := ctx.Value("name")

	// id := r.Context().Value("id")
	// name := r.Context().Value("name")

	fmt.Println("id user context:", id)
	fmt.Println("id user context:", name)

	items := []TransactionItem{}
	totalPrice := 0

	for _, requestItem := range transactionRequest.Items {
		product, err := usecase.ProductRepo.GetProductById(int(requestItem.ProductId))

		if err != nil {
			return nil, fmt.Errorf("Product id not found %d", requestItem.ProductId)
		}

		if product.Deleted_at.IsZero() != true {
			return nil, fmt.Errorf("product already deleted")
		}

		if requestItem.Quantity > product.Stock {
			// return nil, fmt.Errorf("Stock not enough %s", product.Name)
			return nil, err
		}

		price := int(requestItem.Quantity) * product.Price

		item := &TransactionItem {
			ProductId: int(requestItem.ProductId),
			Quantity: requestItem.Quantity,
			Price: price,
		}

		items = append(items, *item)

		totalPrice += price
		product.Stock = product.Stock - requestItem.Quantity

		err = usecase.ProductRepo.EditProduct(int(requestItem.ProductId), product)
		if err != nil {
			return nil, fmt.Errorf("cannot change data")
		}
	}
	// admindata := []AdminData{}
	// a := &AdminData {
	// 	Id: id.(int),
	// 	Name: name.(string),
	// }

	// admindata = append(admindata, *a)

	transaction := &Transaction {
		Timestamp:	time.Now(),
		AdminId:	id.(int),
		Total:		totalPrice,
		Items:		items,
	}

	err := usecase.TransactionRepo.Create(transaction)
	if err != nil {
		return nil, fmt.Errorf("cannot added data")
	}

	newTransaction, err := usecase.TransactionRepo.GetById(transaction.Id)
	if err != nil {
		return nil, fmt.Errorf("error")
	}

	return newTransaction, nil
}

// func (usecase Usecase) CreateTransaction(transactionRequest *CreateTransactionRequest) (*Transaction, error)  {
// 	var transaction Transaction
// 	var total int
// 	for index, item := range transactionRequest.Items {
// 		stringProductId := strconv.Itoa(item.ProductId)
// 	fmt.Println("stok", item.Quantity)

// 		product, err := usecase.ProductRepo.GetProductById(stringProductId)
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println("prod stok", product.Stock)
// 		if item.Quantity > product.Stock {			
// 			return errors.New("stock tidak cukup")
// 		}
// 		total += item.Quantity * product.Price
// 		transactionRequest.Items[index].Price = product.Price
// 		fmt.Println("item price", item.Price)
// 	}
// 		transaction.Timestamp = time.Now()
// 	transaction.Total = total
// 	transaction.Items = transactionRequest.Items
// 	fmt.Println(transactionRequest.Items)
// 	errr := usecase.Repo.CreateTransaction(transaction)
// 	return errr
// }