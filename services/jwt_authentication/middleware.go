package jwt_authentication

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func AuthJWTMiddlewareUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/api/auth/authenticate/user" && r.Method == http.MethodPost {
			next.ServeHTTP(w, r)
			return
		}

		var authorization string
		if authorization = r.Header.Get("Authorization"); authorization == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenString := strings.Replace(authorization, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			log.Println("token = ")
			log.Println(token)
			log.Println("token.Method = ")
			log.Println(token.Method)
			log.Println("token.Method.(*jwt.SigningMethodHMAC) = ")
			log.Println(token.Method.(*jwt.SigningMethodHMAC))

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)

		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})

}

func (aj *AuthJWT) AuthJWTMiddlewareAuthor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/api/auth/authenticate/author" && r.Method == http.MethodPost {
			next.ServeHTTP(w, r)
			return
		}

		var authorization string
		if authorization = r.Header.Get("Authorization"); authorization == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenString := strings.Replace(authorization, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			log.Println("token = ")
			log.Println(token)
			log.Println("token.Method = ")
			log.Println(token.Method)
			log.Println("token.Method.(*jwt.SigningMethodHMAC) = ")
			log.Println(token.Method.(*jwt.SigningMethodHMAC))

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method == JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)

		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
