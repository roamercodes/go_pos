package products

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) CreateTransactionItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var transactionItem TransactionItem
	err := json.NewDecoder(r.Body).Decode(&transactionItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.Usecase.CreateTransactionItem(transactionItem)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func (handler Handler) GetAllTransactionItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	transactionItems, err := handler.Usecase.GetAllTransactionItems()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	hasil, err := json.Marshal(transactionItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func (handler Handler) GetTransactionItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	transactionItem, err := handler.Usecase.GetTransactionItemById(id)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	hasil, err := json.Marshal(transactionItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func (handler Handler) UpdateTransactionItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var transactionItem TransactionItem
	err := json.NewDecoder(r.Body).Decode(&transactionItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = handler.Usecase.UpdateTransactionItemById(id, transactionItem)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func (handler Handler) DeleteTransactionItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	err := handler.Usecase.DeleteTransactionItemById(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("succes"))
}