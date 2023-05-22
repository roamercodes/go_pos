package main

import (
	"fmt"
	"net/http"
	"project-week1/modules/products"
	"project-week1/modules/transactions"
	"project-week1/modules/login"
	"project-week1/middlewares"
	"strconv"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_gofinal?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	productRepo := products.Repository{DB: db}
	productUsecase := products.Usecase{Repo: productRepo}
	productHandler := products.Handler{Usecase: productUsecase}

	transactionRepo := transactions.Repository{DB: db}
	transactionUsecase := transactions.Usecase{TransactionRepo: transactionRepo, ProductRepo: productRepo}
	transactionHandler := transactions.Handler{Usecase: transactionUsecase}

	loginRepo 		:= login.Repository{DB: db}
	loginUsecase 	:= login.Usecase{LoginRepo: loginRepo}
	loginHandler	:= login.Handler{Usecase: loginUsecase}

	r := mux.NewRouter()
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/{id}", middlewares.JwtMiddleware(productHandler.GetProduct)).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}/status", productHandler.SoftDeleteProduct).Methods("PATCH")
	// r.HandleFunc("/products/{id}/status", productHandler.SoftDeleteProduct).Methods("PATCH")
	// r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	r.HandleFunc("/transactions", middlewares.JwtMiddleware(transactionHandler.CreateTransaction)).Methods("POST")
	// r.HandleFunc("/transactions", transactionHandler.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions", transactionHandler.GetAllTransactions).Methods("GET")
	// r.HandleFunc("/transactions/{ids}", transactionHandler.GetTransaction).Methods("GET")
	r.HandleFunc("/transactions/{ids}", middlewares.JwtMiddleware(transactionHandler.GetTransaction)).Methods("GET")

	r.HandleFunc("/login", loginHandler.HandlerLogin).Methods("POST")

	PORT := 8080
	fmt.Println("starting web server at localhost:", PORT)
	if err := http.ListenAndServe(":"+strconv.Itoa(PORT), r); err != nil {
		fmt.Println(err)
		return
	}
}