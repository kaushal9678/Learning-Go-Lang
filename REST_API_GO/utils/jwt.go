package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "your_secret_key"
func GenerateToken(email string, userId int64) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func VerifyToken(token string) (int64,error){
parsedToken, err	:= jwt.Parse(token, func(token *jwt.Token) (interface{}, error){
		_,ok := token.Method.(*jwt.SigningMethodHMAC); if !ok{
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token " + err.Error())
	}
	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims); if !ok || !parsedToken.Valid {
		return 0, errors.New("invalid token claims")
	}
	email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	println("Email:", email)
	println("User ID:", userId)
	return userId,nil
	// You can return the claims or use them as needed
	// return claims, nil		
}
