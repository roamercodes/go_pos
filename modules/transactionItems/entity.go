package transactionsItems

type TransactionItem struct {
	Id				int 	`json:"id"`
	TransactionId	int 	`json:"transaction_id"`
	ProductId		int 	`json:"product_id"`
	Quantity		int		`json:"quantity"`
	Price			int		`json:"price"`
}