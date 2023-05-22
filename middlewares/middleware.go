package middlewares

import (
	"net/http"
	"github.com/golang-jwt/jwt/v4"
	"project-week1/modules/login"
	"strings"
	"context"
	// "fmt"
)

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.ParseWithClaims(tokenString, &login.MyClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(login.JWT_SIGNATURE_KEY), nil
		})
		
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		claims, ok := token.Claims.(*login.MyClaims)

		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token invalid"))
			return
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "idp", claims.Idp)
		ctx = context.WithValue(ctx, "name", claims.Fullname)
		next(w, r.WithContext(ctx))
		// next(w, r)
	}
}