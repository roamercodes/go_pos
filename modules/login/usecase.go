package login

import (
	// "encoding/json"
	"time"
	"github.com/golang-jwt/jwt/v4"
	// "golang.org/x/crypto/bcrypt"
	"fmt"
)

type Usecase struct {
	LoginRepo Repository
}

// type 

func (usecase Usecase) CheckUsernameAndPassword(username, password string) (string, error) {
	user, err := usecase.LoginRepo.CheckUsernameAndPassword(username, password)

	if err != nil {
		return "", err
	}

	fmt.Println(user.Fullname)


	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Idp: user.Id,
		Fullname: user.Fullname,
	}


	// claims := map[string]any{
	// 	// "StandardClaims": jwt.StandardClaims{
	// 	// 	Issuer: APPLICATION_NAME,
	// 	// 	ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
	// 	// },
	// 	"id" : user.Id,
	// 	"name" : user.Fullname,
	// }


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}