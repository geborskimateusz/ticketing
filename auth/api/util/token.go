package util

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	common "github.com/geborskimateusz/ticketing-common"
)

func CreateToken(userid, email string) (string, error) {
	claims := common.CustomClaims{
		Email: email,
		ID:    userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil

}
