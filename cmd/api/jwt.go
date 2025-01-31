package api

import (
	"errors"
	"library/db/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtKey = "libraryAppKey"

func parseJWT(tokenStr string) (*types.UserAuth, error) {
	var user types.UserAuth
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token's signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	//user.ID = claims["id"].(int64)
	user.ID = int64(claims["id"].(float64))
	user.Role = claims["role"].(string)
	expiration := claims["exp"].(float64)
	if time.Now().Unix() > int64(expiration) {
		return nil, errors.New("token has expired")
	}
	return &user, nil

}

// genJWT generates a new JWT token for the given user.
func genJWT(id int64, role string) (string, error) {
	exp := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
