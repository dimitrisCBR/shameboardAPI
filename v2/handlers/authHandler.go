package handlers

import (
	"net/http"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

func GetToken (database *mgo.Database, signingKey []byte) http.HandlerFunc {
	return func ( w http.ResponseWriter, r *http.Request) {
		/* Create the token */
		token := jwt.New(jwt.SigningMethodHS256)

		/* Set token claims */
		claims := make(jwt.MapClaims)
		claims["exp"] =  time.Now().Add(time.Hour * 24).Unix()
		claims["iat"] = time.Now().Unix()
		claims["admin"] = true

		/* Sign the token with our secret */
		tokenString, _ := token.SignedString(signingKey)

		/* Finally, write the token to the browser window */
		w.Write([]byte(tokenString))
	}
}


func AuthMiddleWare(signingKey []byte) *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}