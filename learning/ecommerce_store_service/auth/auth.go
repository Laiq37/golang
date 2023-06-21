package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey []byte

type JWTClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  string `json:"type"`
	jwt.StandardClaims
}

type ValidUser struct {
	Email string
	Type  string
}

func GenerateJWT(name string, email string, userType string) (tokenString string, expiresAt int64, err error) {
	expirationTime := time.Now().AddDate(0, 0, 6)
	claims := &JWTClaim{
		Name:  name,
		Email: email,
		Type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(JwtKey)
	expiresAt = claims.ExpiresAt
	return tokenString, expiresAt, err
}

func Validate(signedToken string) (user ValidUser, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return user, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return user, errors.New("failed to authorize")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return user, errors.New("token expired, authorization required!")
	}
	user = ValidUser{
		Email: claims.Email,
		Type:  claims.Type,
	}
	return user, nil
}
