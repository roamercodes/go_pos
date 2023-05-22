package transactions

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
	"html"
	// "os"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// id := r.Context().Value("id")
	// name := r.Context().Value("name")

	// fmt.Println("id user context:", id)
	// fmt.Println("id user context:", name)

	var createItemRequest CreateTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&createItemRequest)
	if err != nil {
		errorMessage, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	transaction, err := handler.Usecase.CreateTransaction(r.Context(), &createItemRequest)
	if err != nil {
		errorMessage, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func (handler Handler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	transactions, err := handler.Usecase.GetAllTransactions()

	if err != nil {
		errorMessage, _ := json.Marshal(map[string]string{"message": "data not found!"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	_, err = json.Marshal(transactions)
	if err != nil {
		errorMessage, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json!"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Data": transactions,
		"message": "success!",
	})}

func (handler Handler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)

	fmt.Fprintf(w, "%q", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "r.query: %q", r.URL.Query().Get("ids"))

	idd := r.Context().Value("idp")

	fmt.Println(idd)

	id, err := strconv.Atoi(vars["ids"])
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	// id := vars["id"]
	fmt.Println("get by id:", id)
	transaction, err := handler.Usecase.GetTransactionById(id)

	if err != nil {
		errorMessage, _ := json.Marshal(map[string]string{"message": "data not found!"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	_, err = json.Marshal(transaction)
	if err != nil {
		errorMessage, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json!"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Data": transaction,
		"message": "success!",
	})
}
