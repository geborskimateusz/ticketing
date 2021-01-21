package util

import (
	b64 "encoding/base64"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func CreateToken(userid, email string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["email"] = email
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func GenerateJWTtoken(id, email string) (string, error) {
	token, err := CreateToken(id, email)
	if err != nil {
		return "", err
	}

	encodedToken := b64.StdEncoding.EncodeToString([]byte(token.AccessToken))

	return encodedToken, nil
}
