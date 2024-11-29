package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(userId string) (string,error) {
	claims:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub": userId,
		"iss":"gin-crud",
		"exp":time.Now().Add(time.Hour).Unix(),
		"iat":time.Now().Unix(),
	})

	token,err:=claims.SignedString(secretKey)
	if err!=nil{
		return "",err
	}
	
	return token,nil
}
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

