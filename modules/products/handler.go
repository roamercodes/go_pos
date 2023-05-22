package products

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
	"strings"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.Usecase.CreateProduct(product)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func (handler Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	products, err := handler.Usecase.GetAllProducts()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	hasil, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func (handler Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	// id := vars["id"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		messageErr, _ := json.Marshal(map[string]string{"message": "Id not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	} 

	product, err := handler.Usecase.GetProductById(id)

	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
	// hasil, err := json.Marshal(product)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(hasil)
	
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	_, err = json.Marshal(product)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "data found",
		"Data":    product,
	})

	// response := &ResponseAddAndEditData{
	// 	Message: "success",
	// 	Data: product,
	// }
	// json.NewEncoder(w).Encode(response)
}

func (handler Handler) SoftDeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorMessage, _ := json.Marshal(map[string]string{"message": "id not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	var status StatusRequest

	statusRequest := json.NewDecoder(r.Body).Decode(&status)

	if statusRequest != nil {
		errorMessage, _ := json.Marshal(map[string]string{"message": "Failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	// fmt.Println(status)

	if strings.EqualFold(status.Status, "inactive")  {
		err := handler.Usecase.SoftDeleteInactiveProduct(id, status)
		fmt.Println("status inactive")
		if err != nil {
			errorMessage, _ := json.Marshal(map[string]string{"message": "Product already deleted yet"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorMessage)
			return
		}
	} else if strings.EqualFold(status.Status, "active") {
		err := handler.Usecase.SoftDeleteActiveProduct(id, status)
		fmt.Println("status active")
		if err != nil {
			messageErr, _ := json.Marshal(map[string]string{"message": "Product is not deleted yet"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(messageErr)
			return
		}
	}



	// else if active != nil {
	// 	errorMessage, _ := json.Marshal(map[string]string{"message": "Product is not deleted yet"})
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write(errorMessage)
	// 	return
	// }


	product, err := handler.Usecase.GetProductById(id)

	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	// _, err = json.Marshal(product)
	// if err != nil {
	// 	messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write(messageErr)
	// 	return
	// }

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "success",
		"Data":    product,
	})
}

func (handler Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "id not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	var product Product
	product.Id = id

	err = json.NewDecoder(r.Body).Decode(&product)

	// product, err := handler.Usecase.GetProductById(id)


	// if err != nil {
	// 	messageErr, _ := json.Marshal(map[string]string{"message": "Failed to decode json"})
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write(messageErr)
	// 	return
	// }

	// err = handler.Usecase.UpdateProductById(id, &product)
	// if err != nil {
	// 	messageErr, _ := json.Marshal(map[string]string{"message": "Data was not changed"})
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write(messageErr)
	// 	return
	// }

	// // err = handler.Usecase.UpdateProductById(id, product)


	// if err != nil {
	// 	messageErr, _ := json.Marshal(map[string]string{"message": "Data cannot be changed"})
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write(messageErr)
	// 	return
	// }
	
	// response := &ResponseAddAndEditData{
	// 	Message: "Data changed successfully",
	// 	Data: product,
	// }
	// json.NewEncoder(w).Encode(response)

	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
	// w.Write([]byte("succes"))
}

// func (handler Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	err := handler.Usecase.DeleteProductById(id)
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	w.Write([]byte("succes"))
// }