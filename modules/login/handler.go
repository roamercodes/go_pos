package login

import (
	"encoding/json"
	"net/http"
	// "time"
	"fmt"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user User
	// var tok UserResponse

	json.NewDecoder(r.Body).Decode(&user)
	// var result User

	token, err := handler.Usecase.CheckUsernameAndPassword(user.Username, user.Password)
	// fmt.Println(result.Password)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "check invalid username or password"})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	_, err = json.Marshal(token)
	if err != nil {
		errorMessage, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json!"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Data": token,
		"message": "success!",
	})

	fmt.Println("token :", token)
}